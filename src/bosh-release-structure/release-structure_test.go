package releasestructure

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReleaseStructureCheck", func() {
	Context("DirectoryIsAvailable", func() {
		It("Should be Available", func() {
			Expect(IsAvailable("../../bosh-release_test/jobs")).To(BeTrue())
		})
	})
	Context("InvalidFileOrDirectory", func() {
		It("Should thrown an error", func() {
			_, err := IsAvailable("lalala")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(" not exist or is on the wrong place"))

		})
	})
})
