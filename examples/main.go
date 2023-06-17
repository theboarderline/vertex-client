package main

import (
	vertexai "github.com/theboarderline/vertex-client/vertex"
	"log"
)

func main() {
	client := vertexai.NewClient()

	prompt := "Who do most people consider to be the best basketball player of all time"
	maxTokens := 20

	response, err := client.ChatResponse(prompt, maxTokens)
	if err != nil {
		log.Printf("unable to get vertex response: %s", err.Error())
	} else {
		log.Printf("Vertex Response: %s", response)
	}
}
