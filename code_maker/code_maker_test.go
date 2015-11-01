package code_maker_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/trayo/go_mastermind/code_maker"
)

var _ = Describe("CodeMaker", func() {
	It("generates a code", func() {
		code := code_maker.Generate()

		Expect(regexp.MatchString("^[rgby]{4}$", code)).To(BeTrue())
	})

	It("generates a code 4 colors long", func() {
		code := code_maker.Generate()

		Expect(len(code)).To(Equal(4))
	})

	It("generates a new code every time", func() {
		code := code_maker.Generate()
		code = code_maker.Generate()
		code = code_maker.Generate()

		Expect(len(code)).To(Equal(4))
	})

	It("accepts a specific code", func() {
		code := code_maker.Generate("rrrr")

		Expect(code).To(Equal("rrrr"))
	})

	It("ignores other arguments", func() {
		code := code_maker.Generate("yyyy", "bbbb", "gggg")

		Expect(code).To(Equal("yyyy"))
	})
})
