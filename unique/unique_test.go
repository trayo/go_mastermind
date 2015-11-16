package unique_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind/unique"
)

var _ = Describe("removing duplicate elements", func() {
	It("can remove duplicate letters", func() {
		elements := "rygrygb"

		newElements := unique.RemoveDuplicateLetters(elements)

		Expect(newElements).To(Equal("rygb"))
	})

	It("can return an array of a unique letter", func() {
		elements := "rrrrrrr"

		newElements := unique.RemoveDuplicateLetters(elements)

		Expect(newElements).To(Equal("r"))
	})
})
