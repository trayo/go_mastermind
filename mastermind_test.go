package main_test

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/stretchr/testify/mock"

	"github.com/trayo/go_mastermind"
	"github.com/trayo/go_mastermind/print"
	"github.com/trayo/go_mastermind/test/mocks"
)

var _ = Describe("Main CLI interaction", func() {

	var (
		stdin           *bytes.Buffer
		stdinReader     *bufio.Reader
		buffer          *gbytes.Buffer
		printer         print.Printer
		mockGamer       *mocks.Gamer
		commandSequence = func(args ...string) {
			stdin.WriteString(strings.Join(args, "\n"))
		}
	)

	BeforeEach(func() {
		mockGamer = &mocks.Gamer{}

		stdin = &bytes.Buffer{}
		stdinReader = bufio.NewReader(stdin)

		buffer = gbytes.NewBuffer()

		printer = print.NewPrinter(buffer)
	})

	AfterEach(func() {
		Expect(mockGamer.AssertExpectations(GinkgoT())).To(BeTrue())
	})

	Context("when providing valid input", func() {
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

			mockGamer.On("Play", mock.Anything).Times(1)

			main.Run(mockGamer, stdinReader, printer)
		})
	})

	Context("entering the cheat code", func() {
		It("can accept a code", func() {
			secret := os.Getenv("MASTERMIND_SECRET")
			code := "rrrr"
			commandSequence(secret, code, "q")

			mockGamer.On("Play", mock.Anything).Times(1)

			main.Run(mockGamer, stdinReader, printer)
		})
	})

	Context("when providing invalid input", func() {
		It("will ask for valid input", func() {
			commandSequence("this wont work", "q")

			main.Run(mockGamer, stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Unknown command"))
			Eventually(buffer).Should(gbytes.Say("Valid commands are"))
			Eventually(buffer).Should(gbytes.Say("Would you like to play"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})
	})
})
