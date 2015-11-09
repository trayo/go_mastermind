package check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind/check"
)

var (
	code   string
	guess  string
	answer string
)

var _ = Describe("checking postions", func() {

	It("can check for 0 correct positions", func() {
		code = "gggg"
		guess = "bbbb"

		answer = check.Positions(guess, code)

		Expect(answer).To(Equal("0"))
	})

	It("can check for 1 correct position", func() {
		code = "grgg"
		guess = "rrrr"

		answer = check.Positions(guess, code)

		Expect(answer).To(Equal("1"))
	})

	It("can check for 2 correct positions", func() {
		code = "ybyb"
		guess = "yyyy"

		answer = check.Positions(guess, code)

		Expect(answer).To(Equal("2"))
	})

	It("can check for 3 correct positions", func() {
		code = "gbgg"
		guess = "gggg"

		answer = check.Positions(guess, code)

		Expect(answer).To(Equal("3"))
	})

	It("can check for all correct positions", func() {
		code = "gggg"
		guess = "gggg"

		answer = check.Positions(guess, code)

		Expect(answer).To(Equal("4"))
	})
})
