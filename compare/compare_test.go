package compare_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind/compare"
)

var _ = Describe("comparing guesses and codes", func() {

	var comp *compare.Compare

	BeforeEach(func() {
		comp = &compare.Compare{}
	})

	It("has correctPositions and correctColors", func() {
		comp = &compare.Compare{
			CorrectPositions: "0",
			CorrectColors:    "0",
		}

		Expect(comp.CorrectPositions).To(Equal("0"))
		Expect(comp.CorrectColors).To(Equal("0"))
	})

	It("compares guesses upon creation", func() {
		code := "rgrg"
		guess := "gggg"

		comp = compare.NewCompare(guess, code)

		Expect(comp.CorrectPositions).To(Equal("2"))
		Expect(comp.CorrectColors).To(Equal("2"))
	})

	It("can compare new guesses", func() {
		code := "ybgr"
		guess := "rybg"
		comp = compare.NewCompare(guess, code)

		comp.Guess("ybgr")

		Expect(comp.CorrectPositions).To(Equal("4"))
		Expect(comp.CorrectColors).To(Equal("4"))
	})

	It("can use the code package to make a random code", func() {
		guess := "rybg"

		comp = compare.NewCompare(guess)
		Expect(regexp.MatchString("^[rgby]{4}$", comp.Code)).To(BeTrue())
	})
})
