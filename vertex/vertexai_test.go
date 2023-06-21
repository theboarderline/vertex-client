package vertexai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/vertex-client/vertex"
)

var _ = Describe("Vertexai", func() {
	It("can get a comment using vertex AI", func() {
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
})
