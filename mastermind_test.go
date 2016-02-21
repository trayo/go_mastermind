package main_test

import (
	"bufio"
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/trayo/go_mastermind"
	"github.com/trayo/go_mastermind/print"
	mocks "github.com/trayo/go_mastermind/test/mocks"
)

var _ = Describe("Main CLI interaction", func() {

	var (
		stdin           *bytes.Buffer
		stdinReader     *bufio.Reader
		buffer          *gbytes.Buffer
		printer         print.Printer
		mockGamer       *mocks.MockGamer
		commandSequence = func(args ...string) {
			stdin.WriteString(strings.Join(args, "\n"))
		}
	)

	BeforeEach(func() {
		mockGamer = mocks.NewMockGamer(mockCtrl)

		stdin = &bytes.Buffer{}
		stdinReader = bufio.NewReader(stdin)

		buffer = gbytes.NewBuffer()

		printer = print.NewPrinter(buffer)
	})

	Context("when providing input", func() {
		It("can quit the game", func() {
			commandSequence("q")

			main.Run(mockGamer, stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Would you like to play"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})

		It("can visit the instructions page", func() {
			commandSequence("i", "q")

			main.Run(mockGamer, stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Would you like to play"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})

		It("can visit the instructions page multiple times", func() {
			commandSequence("i", "i", "q")

			main.Run(mockGamer, stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Would you like to play"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})

		It("can start a game", func() {
			commandSequence("p", "q")

			mockGamer.EXPECT().Play()

			main.Run(mockGamer, stdinReader, printer)
		})
	})
})
