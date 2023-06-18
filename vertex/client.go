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
	ProjectID   string
	AccessToken string
	Response    string
	httpClient  *resty.Client
}

func NewClient(projectIDs ...string) *Client {

	var projectID string
	if len(projectID) == 0 {
		projectID = os.Getenv("GCP_PROJECT_ID")
		if projectID == "" {
			log.Fatal().Err(errors.New("GCP_PROJECT_ID not set"))
		}
	} else {
		projectID = projectIDs[0]
	}

	accessToken, err := getAccessToken()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get access token")
	}

	return &Client{
		ProjectID:   projectID,
		AccessToken: accessToken,
		httpClient:  resty.New(),
	}
}

func (c Client) FormatVertexAPIURL(modelID string) string {
	return fmt.Sprintf("https://%s/v1/projects/%s/locations/us-central1/publishers/google/models/%s:predict", API_ENDPOINT, c.ProjectID, modelID)
}

type ChatRequest struct {
	Debug      bool
	Prompt     string
	ModelID    string
	Parameters []Parameter
}

func (c Client) ChatResponse(req ChatRequest) (string, error) {

	if req.ModelID == "" {
		req.ModelID = BISON_MODEL_ID
	}

	if len(req.Parameters) == 0 {
		req.Parameters = []Parameter{DEFAULT_PARAMETER}
	}

	request := Request{
		Instances:  []Instance{{Prefix: req.Prompt}},
		Parameters: req.Parameters,
	}
	url := c.FormatVertexAPIURL(req.ModelID)

	if req.Debug {
		log.Debug().Msgf("vertexai api url: %s", url)
		log.Debug().Msgf("vertexai request: %+v", request)
	}

	response, err := c.httpClient.R().
		SetAuthToken(c.AccessToken).
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

	if len(res.Predictions) == 0 {
		err = errors.New("no predictions returned")
		log.Err(err)
		return "", err
	}

	if req.Debug {
		log.Debug().Msgf("vertexai response: %+v", res)
	}

	return res.Predictions[0].Content, nil
}
