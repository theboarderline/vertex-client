package vertexai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/vertex-client/vertex"
	"os"
)

var _ = Describe("Vertexai", func() {

	It("can return an error if no project is specified", func() {
		oldProjectID := os.Getenv("GCP_PROJECT_ID")
		os.Setenv("GCP_PROJECT_ID", "")

		_, err := vertexai.NewClient(vertexai.ClientConfig{Debug: true})
		Expect(err).To(HaveOccurred())

		os.Setenv("GCP_PROJECT_ID", oldProjectID)
	})

	It("can get a comment using vertex AI code model", func() {
		client, err := vertexai.NewClient(vertexai.ClientConfig{})
		Expect(err).NotTo(HaveOccurred())

		req := vertexai.ChatRequest{
			Prompt:  "Who is the best basketball player of all time?",
			ModelID: vertexai.BISON_CODE_MODEL_ID,
		}
		response, err := client.ChatResponse(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeEmpty())
	})

	It("can get a comment using vertex AI text model", func() {
		client, err := vertexai.NewClient(vertexai.ClientConfig{Debug: true})
		Expect(err).NotTo(HaveOccurred())

		req := vertexai.ChatRequest{
			Prompt:    "Come up with a funny joke",
			ModelID:   vertexai.BISON_TEXT_MODEL_ID,
			Instances: []vertexai.Instance{{Prefix: "Q: What do you call a fake noodle?"}},
		}
		response, err := client.ChatResponse(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeEmpty())
	})

})
