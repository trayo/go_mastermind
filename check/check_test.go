package check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/trayo/go_mastermind/check"
)

var (
	code   string
	guess  string
	answer int
	comp   *Check
)

var _ = Describe("Check", func() {

	BeforeEach(func() {
		comp = &Check{}
	})

	It("can be initialized empty", func() {
		comp = NewCheck()

		Expect(comp.CorrectPositions).To(Equal(0))
		Expect(comp.CorrectColors).To(Equal(0))
	})

	It("can compare new guesses", func() {
		code := "ybgr"
		comp = NewCheck(code)

		comp.Guess(code)

		Expect(comp.CorrectPositions).To(Equal(4))
		Expect(comp.CorrectColors).To(Equal(4))
	})

	Context("checking positions", func() {
		It("can check for 0 correct positions", func() {
			code = "gggg"
			guess = "bbbb"

			answer = Positions(guess, code)

			Expect(answer).To(Equal(0))
		})

		It("can check for 1 correct position", func() {
			code = "grgg"
			guess = "rrrr"

			answer = Positions(guess, code)

			Expect(answer).To(Equal(1))
		})

		It("can check for 2 correct positions", func() {
			code = "ybyb"
			guess = "yyyy"

			answer = Positions(guess, code)

			Expect(answer).To(Equal(2))
		})

		It("can check for 3 correct positions", func() {
			code = "gbgg"
			guess = "gggg"

			answer = Positions(guess, code)

			Expect(answer).To(Equal(3))
		})

		It("can check for all correct positions", func() {
			code = "gggg"
			guess = "gggg"

			answer = Positions(guess, code)

			Expect(answer).To(Equal(4))
		})
	})

	Context("checking colors", func() {
		It("can check for 0 correct colors", func() {
			code = "gggg"
			guess = "rrrr"

			answer = Colors(guess, code)

			Expect(answer).To(Equal(0))
		})

		It("can check for 1 correct color", func() {
			code = "bggg"
			guess = "bbbb"

			answer = Colors(guess, code)

			Expect(answer).To(Equal(1))
		})

		It("can check for 2 correct colors", func() {
			code = "rbrb"
			guess = "rrrr"

			answer = Colors(guess, code)

			Expect(answer).To(Equal(2))
		})

		It("can check for 3 correct colors", func() {
			code = "ggyg"
			guess = "gggg"

			answer = Colors(guess, code)

			Expect(answer).To(Equal(3))
		})

		It("can check for all correct colors", func() {
			code = "rybg"
			guess = "gbyr"

			answer = Colors(guess, code)

			Expect(answer).To(Equal(4))
		})
	})
})
