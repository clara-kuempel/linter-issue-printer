package executable_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExecutable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Executable Suite")
}
