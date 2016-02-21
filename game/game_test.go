package game_test

import (
	"bufio"
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"

	. "github.com/trayo/go_mastermind/game"
)

var _ = Describe("playing a game", func() {

	var (
		stdin       *bytes.Buffer
		stdinReader *bufio.Reader
		buffer      *gbytes.Buffer
		printer     print.Printer
		gamer       I.Gamer
		// commandSequence = func(args ...string) {
		// 	stdin.WriteString(strings.Join(args, "\n"))
		// }
	)

	BeforeEach(func() {
		stdin = &bytes.Buffer{}
		stdinReader = bufio.NewReader(stdin)

		buffer = gbytes.NewBuffer()

		printer = print.NewPrinter(buffer)

		gamer = NewGamer(stdinReader, printer)
	})

	Context("starting a game", func() {
		It("says it generates a code", func() {
			gamer.Play()

			Eventually(buffer).Should(gbytes.Say("I have generated a random code"))
		})

		It("asks for a guess", func() {
			gamer.Play()

			Eventually(buffer).Should(gbytes.Say("Enter a guess"))
		})
	})

	// It("can make a guess", func() {
	// 	commandSequence("")
	// 	Expect(true).To(Equal(true))
	// })
})
