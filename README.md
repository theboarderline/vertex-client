
# Google Vertex AI API Client

## Prerequisites
- [Golang](https://go.dev/doc/install)
- [Gcloud](https://cloud.google.com/sdk/docs/install)

Since the Vertex AI API does not allow API Keys,
the API Client will automatically use the credentials
from the environment to create an access token to be used
in the request.

Access tokens are short-lived, so a new client should be created for each request.

### Usage

Authenticate with user GCP credentials

```bash
gcloud auth application-default login
```

### IAM
- Requires Vertex AI User Role - [Vertex AI IAM](https://cloud.google.com/vertex-ai/docs/general/access-control)

### Example

```go
package main

import (
	vertexai "github.com/theboarderline/vertex-client/vertex"
	"log"
)

func main() {
	client := vertexai.NewClient() // uses environment variable GCP_PROJECT_ID
	
	// or

	client = vertexai.NewClient("PROJECT_ID") // pass in project id

	prompt := "Who do most people consider to be the best basketball player of all time"
	maxTokens := 20

	response, err := client.ChatResponse(prompt, maxTokens)
	if err != nil {
		log.Printf("unable to get vertex response: %s", err.Error())
	} else {
		log.Printf("Vertex Response: %s", response)
	}
}

```
