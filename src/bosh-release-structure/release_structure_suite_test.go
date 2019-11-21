package releasestructure

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReleaseStructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReleaseStructure Suite")
}
