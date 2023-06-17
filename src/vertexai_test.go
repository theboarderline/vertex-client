package vertexai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/rambler/api/src/bots/memecity/commentator"
	"github.com/theboarderline/rambler/api/src/pkg/vertexai"
)

var _ = Describe("Vertexai", func() {
	It("can get a comment for repeat bot using vertex AI", func() {
		client := vertexai.NewClient()
		response, err := client.ChatResponse(commentator.AIPrompt, 20)
		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeEmpty())
	})
})
