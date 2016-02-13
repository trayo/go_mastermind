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
			s := "q"

			Expect(WantsToQuit(s)).To(BeTrue())
		})

		It("can check for quit", func() {
			s := "quit"

			Expect(WantsToQuit(s)).To(BeTrue())
		})
	})

	Context("instructions", func() {
		It("can check for i", func() {
			s := "i"

			Expect(WantsInstructions(s)).To(BeTrue())
		})

		It("can check for instructions", func() {
			s := "instructions"

			Expect(WantsInstructions(s)).To(BeTrue())
		})
	})

	Context("play", func() {
		It("can check for p", func() {
			s := "p"

			Expect(WantsToPlay(s)).To(BeTrue())
		})

		It("can check for play", func() {
			s := "play"

			Expect(WantsToPlay(s)).To(BeTrue())
		})
	})
})
