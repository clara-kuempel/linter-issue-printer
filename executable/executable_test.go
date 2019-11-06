package executable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "clara.org/lip"
)

var _ = Describe("ExecutableCheck", func() {
	Context("IsExecutable", func() {
		It("it should be executable", func() {
			Expect(IsExecutable("permission_checker")).To(BeTrue())
		})
	})
	Context("when invalid file passed", func() {
		It("should thrown an error", func() {
			_, err := IsExecutable("lalal")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("does not exist"))
		})
	})
})
