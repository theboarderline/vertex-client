package main

import (
	vertexai "github.com/theboarderline/vertex-client/vertex"
	"log"
)

func main() {
	client := vertexai.NewClient()
	response, err := client.ChatResponse("Who do most people consider to be the best basketball player of all time", 20)
	if err != nil {
		log.Printf("unable to get vertex response: %s", err.Error())
	} else {
		log.Printf("Vertex Response: %s", response)
	}
}
