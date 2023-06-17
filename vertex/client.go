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
	} else {
		projectID = projectIDs[0]
	}

	accessToken, err := getAccessToken()
	if err != nil {
		log.Err(err).Msg("could not get access token")
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

func (c Client) ChatResponse(prompt string, maxTokens int, modelIDs ...string) (string, error) {

	var modelID string
	if len(modelID) == 0 {
		modelID = BISON_MODEL_ID
	} else {
		modelID = modelIDs[0]
	}

	request := Request{
		Instances: []Instance{{Prefix: prompt}},
		Parameters: []Parameter{
			{
				MaxOutputTokens: maxTokens,
				Temperature:     1,
				TopP:            0,
				TopK:            40,
			},
		},
	}

	response, err := c.httpClient.R().
		SetAuthToken(c.AccessToken).
		SetBody(request).
		Post(c.FormatVertexAPIURL(modelID))

	if err != nil {
		log.Err(err).Msg("could not get messages from groupme api")
		return "", err
	}

	var res Response
	if err = json.Unmarshal(response.Body(), &res); err != nil {
		log.Err(err).Msg("could not unmarshal response from groupme api")
		return "", err
	}

	if len(res.Predictions) == 0 {
		err = errors.New("no predictions returned")
		log.Err(err)
		return "", err
	}

	return res.Predictions[0].Content, nil
}
