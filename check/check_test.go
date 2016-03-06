package check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/trayo/go_mastermind/check"
)

var (
	code    string
	guess   string
	answer  int
	checker *Check
)

var _ = Describe("Check", func() {

	BeforeEach(func() {
		checker = &Check{}
	})

	It("can be initialized empty", func() {
		checker = NewCheck()

		Expect(checker.CorrectPositions).To(Equal(0))
		Expect(checker.CorrectColors).To(Equal(0))
	})

	It("can check the right code", func() {
		code := "ybgr"
		checker = NewCheck(code)

		checker.Guess(code)

		Expect(checker.CorrectPositions).To(Equal(4))
		Expect(checker.CorrectColors).To(Equal(4))
		Expect(checker.Won()).To(BeTrue())
	})

	Context("checking positions", func() {
		It("can check for 0 correct positions", func() {
			code = "gggg"
			guess = "bbbb"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectPositions).To(Equal(0))
		})

		It("can check for 1 correct position", func() {
			code = "grgg"
			guess = "rrrr"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectPositions).To(Equal(1))
		})

		It("can check for 2 correct positions", func() {
			code = "ybyb"
			guess = "yyyy"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectPositions).To(Equal(2))
		})

		It("can check for 3 correct positions", func() {
			code = "gbgg"
			guess = "gggg"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectPositions).To(Equal(3))
		})

		It("can check for 4 correct positions", func() {
			code = "gbrr"
			guess = "gbrr"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectPositions).To(Equal(4))
		})
	})

	Context("checking colors", func() {
		It("can check for 0 correct colors", func() {
			code = "gggg"
			guess = "rrrr"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectColors).To(Equal(0))
		})

		It("can check for 1 correct color", func() {
			code = "bggg"
			guess = "bbbb"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectColors).To(Equal(1))
		})

		It("can check for 1 correct colors", func() {
			code = "rrbb"
			guess = "rrrr"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectColors).To(Equal(1))
		})

		It("can check for 2 correct colors", func() {
			code = "ggyb"
			guess = "bggr"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectColors).To(Equal(2))
		})

		It("counts duplicate colors as one element", func() {
			code = "rgyg"
			guess = "ygrg"
			checker = NewCheck(code)

			checker.Guess(guess)

			Expect(checker.CorrectColors).To(Equal(3))
		})
	})
})
