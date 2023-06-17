package vertexai

const (
	API_ENDPOINT   = "us-central1-aiplatform.googleapis.com"
	BISON_MODEL_ID = "code-bison@001"
)

var (
	DEFAULT_PARAMETER = Parameter{
		MaxOutputTokens: 100,
		Temperature:     1,
		TopP:            0,
		TopK:            40,
	}
)
