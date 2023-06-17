package vertexai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/vertex-client/vertex"
)

var _ = Describe("Vertexai", func() {
	It("can get a comment using vertex AI", func() {
		client := vertexai.NewClient()
		req := vertexai.ChatRequest{
			Prompt:  "I am a bot",
			ModelID: vertexai.BISON_MODEL_ID,
		}
		response, err := client.ChatResponse(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeEmpty())
	})
})
