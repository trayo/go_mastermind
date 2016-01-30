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
)

var _ = Describe("Main CLI interaction", func() {

	var (
		stdin       *bytes.Buffer
		stdinReader *bufio.Reader
		buffer      *gbytes.Buffer
		printer     print.Printer
	)

	BeforeEach(func() {
		stdin = &bytes.Buffer{}
		stdinReader = bufio.NewReader(stdin)

		buffer = gbytes.NewBuffer()

		printer = print.NewPrinter(buffer)
	})

	Context("when providing input", func() {
		It("can quit the game", func() {
			commandSequence(stdin, "q")

			main.Run(stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Welcome"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})

		It("can visit the instructions page", func() {
			commandSequence(stdin, "i", "q")

			main.Run(stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Welcome"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})

		It("can visit the instructions page multiple times", func() {
			commandSequence(stdin, "i", "i", "q")

			main.Run(stdinReader, printer)

			Eventually(buffer).Should(gbytes.Say("Welcome"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("generate a sequence"))
			Eventually(buffer).Should(gbytes.Say("byeee"))
		})
	})
})

func commandSequence(stdin *bytes.Buffer, args ...string) {
	stdin.WriteString(strings.Join(args, "\n"))
}
