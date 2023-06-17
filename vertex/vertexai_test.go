package vertexai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/vertex-client/vertex"
)

var _ = Describe("Vertexai", func() {
	It("can get a comment for repeat bot using vertex AI", func() {
		client := vertexai.NewClient()
		response, err := client.ChatResponse("Who do most people consider to be the best basketball player of all time", 20)
		Expect(err).NotTo(HaveOccurred())
		Expect(response).NotTo(BeEmpty())
	})
})
