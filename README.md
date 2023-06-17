
# Google Vertex AI API Client

## Prerequisites
- [Golang](https://go.dev/doc/install)
- [Gcloud](https://cloud.google.com/sdk/docs/install)

Since the Vertex AI API does not allow API Keys,
the API Client will automatically use the credentials
from the environment to create an access token to be used
in the request.

### Usage

Authenticate with user GCP credentials

```bash
gcloud auth application-default login
```

### Example

```go
package main

import (
	"github.com/rs/zerolog/log"
	vertexai "github.com/theboarderline/vertex-client/vertex"
)

func main() {
	client := vertexai.NewClient()
	response, err := client.ChatResponse("Who do most people consider to be the best basketball player of all time", 20)
	if err != nil {
		log.Err(err).Msg("unable to get vertex response")
	} else {
		log.Info().Msgf("Vertex Response: %s", response)
	}
}
```
