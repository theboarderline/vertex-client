package vertexai_test

import (
	"github.com/theboarderline/rambler/api/src/config"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVertexai(t *testing.T) {
	RegisterFailHandler(Fail)
	config.Init()
	RunSpecs(t, "Vertexai Suite")
}
