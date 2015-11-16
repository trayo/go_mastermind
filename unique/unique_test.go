package unique_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind/unique"
)

var _ = Describe("removing duplicate elements", func() {
	It("can remove duplicate letters", func() {
		elements := []string{"r", "y", "g", "r", "y", "g", "b"}

		newElements := unique.RemoveDuplicateStrings(elements)

		Expect(newElements).To(Equal([]string{"r", "y", "g", "b"}))
	})

	It("can return an array of a unique letter", func() {
		elements := []string{"r", "r", "r", "r", "r", "r", "r"}

		newElements := unique.RemoveDuplicateStrings(elements)

		Expect(newElements).To(Equal([]string{"r"}))
	})
})
