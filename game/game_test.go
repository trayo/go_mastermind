package game_test

import (
	"bufio"
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"

	. "github.com/trayo/go_mastermind/game"
)

var _ = Describe("playing a game", func() {

	var (
		stdin           *bytes.Buffer
		stdinReader     *bufio.Reader
		buffer          *gbytes.Buffer
		printer         print.Printer
		gamer           I.Gamer
		commandSequence = func(args ...string) {
			stdin.WriteString(strings.Join(args, "\n"))
		}
	)

	BeforeEach(func() {
		stdin = &bytes.Buffer{}
		stdinReader = bufio.NewReader(stdin)

		buffer = gbytes.NewBuffer()

		printer = print.NewPrinter(buffer)

		gamer = NewGamer(stdinReader, printer)
	})

	Context("with invalid input", func() {
		It("will ask for valid input", func() {
			commandSequence("hoohah", "this wont work", "q")

			gamer.Play()

			Eventually(buffer).Should(gbytes.Say("Unknown command"))
			Eventually(buffer).Should(gbytes.Say("Valid commands are"))
			Eventually(buffer).Should(gbytes.Say("Unknown command"))
			Eventually(buffer).Should(gbytes.Say("Valid commands are"))
			Eventually(buffer).Should(gbytes.Say("Thanks for playing"))
		})
	})

	Context("with valid input", func() {
		It("starts up", func() {
			commandSequence("q")

			gamer.Play()

			Eventually(buffer).Should(gbytes.Say("I have generated a random code"))
			Eventually(buffer).Should(gbytes.Say("Enter a guess"))
			Eventually(buffer).Should(gbytes.Say("Thanks for playing"))
		})

		It("can guess the correct code", func() {
			code := "rrrr"
			commandSequence("rrrr", "q")

			gamer.Play(code)

			Eventually(buffer).Should(gbytes.Say("I have generated a random code"))
			Eventually(buffer).Should(gbytes.Say("Enter a guess"))
			Eventually(buffer).Should(gbytes.Say("You have guessed the code!"))
			Eventually(buffer).Should(gbytes.Say("Thanks for playing"))
		})

		It("says correct colors and correct positions", func() {
			code := "rrrr"
			commandSequence("gggg", "rrrr", "q")

			gamer.Play(code)

			Eventually(buffer).Should(gbytes.Say("0 correct colors"))
			Eventually(buffer).Should(gbytes.Say("0 correct positions"))
		})

		It("says 3 correct colors and 1 correct positions", func() {
			code := "ygyb"
			commandSequence("gbyr", "q")

			gamer.Play(code)

			Eventually(buffer).Should(gbytes.Say("3 correct colors"))
			Eventually(buffer).Should(gbytes.Say("1 correct positions"))
		})

		It("says how many guesses it took", func() {
			code := "rrrr"
			commandSequence("gggg", "gggg", "rrrr", "q")

			gamer.Play(code)

			Eventually(buffer).Should(gbytes.Say("3 guesses"))
		})

		It("says guess for one guess", func() {
			code := "rrrr"
			commandSequence("rrrr", "q")

			gamer.Play(code)

			Eventually(buffer).Should(gbytes.Say("1 guess"))
		})
	})
})
