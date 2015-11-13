package code_maker_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind/code_maker"
)

var _ = Describe("code maker", func() {
	var (
		codeMaker code_maker.CodeMaker
		code      string
	)

	BeforeEach(func() {
		codeMaker = *code_maker.NewCodeMaker()
	})

	It("generates a code", func() {
		code = codeMaker.Generate()

		Expect(regexp.MatchString("^[rgby]{4}$", code)).To(BeTrue())
	})

	It("generates a code 4 colors long", func() {
		code = codeMaker.Generate()

		Expect(len(code)).To(Equal(4))
	})

	It("generates a new code every time", func() {
		code = codeMaker.Generate()
		code = codeMaker.Generate()
		code = codeMaker.Generate()

		Expect(len(code)).To(Equal(4))
	})

	It("accepts a specific code", func() {
		code = codeMaker.Generate("rrrr")

		Expect(code).To(Equal("rrrr"))
	})

	It("ignores extra arguments", func() {
		code = codeMaker.Generate("yyyy", "bbbb", "gggg")

		Expect(code).To(Equal("yyyy"))
	})
})
