package vertexai

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"os"
)

type Client struct {
	Response    string
	projectID   string
	accessToken string
	debug       bool
	httpClient  *resty.Client
}

type ClientConfig struct {
	ProjectID string
	Debug     bool
}

func NewClient(config ClientConfig) (*Client, error) {

	if config.ProjectID == "" {
		config.ProjectID = os.Getenv("GCP_PROJECT_ID")
		if config.ProjectID == "" {
			err := errors.New("GCP_PROJECT_ID is empty")
			log.Error().Err(err)
			return nil, err
		}
	}

	accessToken, err := getAccessToken()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &Client{
		debug:       config.Debug,
		projectID:   config.ProjectID,
		accessToken: accessToken,
		httpClient:  resty.New(),
	}, nil
}

func (c Client) FormatVertexAPIURL(modelID string) string {
	return fmt.Sprintf("https://%s/v1/projects/%s/locations/us-central1/publishers/google/models/%s:predict", API_ENDPOINT, c.projectID, modelID)
}

type ChatRequest struct {
	Debug      bool
	Prompt     string
	ModelID    string
	Instances  []Instance
	Parameters []Parameter
}

func (c Client) ChatResponse(req ChatRequest) (string, error) {

	if req.ModelID == "" {
		c.debugMsgf("model id is empty, using default model id: %s", BISON_TEXT_MODEL_ID)
		req.ModelID = BISON_CODE_MODEL_ID
	}

	if len(req.Instances) == 0 {
		if req.Prompt == "" {
			err := errors.New("instances and prompt cannot be empty")
			log.Error().Err(err)
			return "", err
		}

		c.debugMsg("instances is empty, using prompt as instance")
		req.Instances = []Instance{{Prefix: req.Prompt}}
	}

	if len(req.Parameters) == 0 {
		req.Parameters = []Parameter{DEFAULT_PARAMETER}

		c.debugMsg("parameters is empty, using default parameter")
	}

	url := c.FormatVertexAPIURL(req.ModelID)
	c.debugMsgf("vertex api url: %s", url)

	request := Request{
		Instances:  req.Instances,
		Parameters: req.Parameters,
	}
	c.debugMsgf("vertex request: %+v", request)

	response, err := c.httpClient.R().
		SetAuthToken(c.accessToken).
		SetBody(request).
		Post(url)

	if err != nil {
		log.Err(err).Msg("could not get response from vertex api")
		return "", err
	}

	var res Response
	if err = json.Unmarshal(response.Body(), &res); err != nil {
		log.Err(err).Msg("could not unmarshal response from vertex ai api")
		return "", err
	}

	if res.Error.Code != 0 {
		msg := fmt.Sprintf("%d %s %s", res.Error.Code, res.Error.Status, res.Error.Message)
		err = errors.New(msg)
		log.Err(err)
		return "", err
	}

	if len(res.Predictions) == 0 {
		err = errors.New("no predictions returned")
		log.Err(err)
		return "", err
	}

	c.debugMsgf("vertexai response: %+v", res)

	return res.Predictions[0].Content, nil
}
