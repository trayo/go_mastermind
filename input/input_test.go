package input_test

import (
	"bufio"
	"bytes"

	. "github.com/trayo/go_mastermind/input"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input", func() {
	var (
		buffer *bytes.Buffer
		reader *bufio.Reader
	)

	BeforeEach(func() {
		buffer = &bytes.Buffer{}
		reader = bufio.NewReader(buffer)
	})

	It("can get input from a reader", func() {
		buffer.WriteString("it's full of stars")

		out := GetInput(reader)

		Expect(out).To(Equal("it's full of stars"))
	})

	Context("quit", func() {
		It("can check for q", func() {
			input := "q"

			Expect(WantsToQuit(input)).To(BeTrue())
		})

		It("can check for quit", func() {
			input := "quit"

			Expect(WantsToQuit(input)).To(BeTrue())
		})
	})

	Context("instructions", func() {
		It("can check for i", func() {
			input := "i"

			Expect(WantsInstructions(input)).To(BeTrue())
		})

		It("can check for instructions", func() {
			input := "instructions"

			Expect(WantsInstructions(input)).To(BeTrue())
		})
	})

	Context("play", func() {
		It("can check for p", func() {
			input := "p"

			Expect(WantsToPlay(input)).To(BeTrue())
		})

		It("can check for play", func() {
			input := "play"

			Expect(WantsToPlay(input)).To(BeTrue())
		})
	})

	Context("checking for valid game input", func() {
		It("can check correct length and colors", func() {
			input := "rgby"

			Expect(Valid(input)).To(BeTrue())
		})

		It("can check invalid colors", func() {
			input := "nope"

			Expect(Valid(input)).To(BeFalse())
		})

		It("can check input that's too long", func() {
			input := "thisistoolong"

			Expect(Valid(input)).To(BeFalse())
		})

		It("can check input that's too short", func() {
			input := "no"

			Expect(Valid(input)).To(BeFalse())
		})
	})
})
