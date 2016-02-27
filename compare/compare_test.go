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
			CorrectPositions: 0,
			CorrectColors:    0,
		}

		Expect(comp.CorrectPositions).To(Equal(0))
		Expect(comp.CorrectColors).To(Equal(0))
	})

	It("can compare new guesses", func() {
		var code = "ybgr"
		comp = compare.NewCompare(code)

		comp.Guess(code)

		Expect(comp.CorrectPositions).To(Equal(4))
		Expect(comp.CorrectColors).To(Equal(4))
	})

	It("can use the code package to make a random code", func() {
		guess := "rybg"

		comp = compare.NewCompare(guess)

		Expect(regexp.MatchString("^[rgby]{4}$", comp.Code)).To(BeTrue())
	})

	It("can be initialized empty", func() {
		comp = compare.NewCompare()

		Expect(comp.CorrectPositions).To(Equal(0))
		Expect(comp.CorrectColors).To(Equal(0))
	})
})
