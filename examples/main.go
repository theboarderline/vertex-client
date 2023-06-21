package main

import (
	vertexai "github.com/theboarderline/vertex-client/vertex"
	"log"
)

func main() {
	client, err := vertexai.NewClient(vertexai.ClientConfig{Debug: true})
	if err != nil {
		log.Printf("unable to create vertex client: %s", err.Error())
		return
	}

	// code model
	request := vertexai.ChatRequest{
		Debug:   true,
		ModelID: vertexai.BISON_CODE_MODEL_ID,
		Prompt:  "Who do most people consider to be the best basketball player of all time",
	}

	response, err := client.ChatResponse(request)
	if err != nil {
		log.Printf("unable to get vertex response: %s", err.Error())
	} else {
		log.Printf("Vertex Response: %s", response)
	}

	// text model
	request = vertexai.ChatRequest{
		Debug:   true,
		ModelID: vertexai.BISON_TEXT_MODEL_ID,
		Instances: []vertexai.Instance{
			{
				Content: "Who is the best basketball player of all time?",
			},
		},
	}

	response, err = client.ChatResponse(request)
	if err != nil {
		log.Printf("unable to get vertex response: %s", err.Error())
	} else {
		log.Printf("Vertex Response: %s", response)
	}
}
