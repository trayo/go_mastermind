package main_test

import (
	"bufio"
	"bytes"

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

	It("can quit the game", func() {
		stdin.WriteString("q\n")

		main.Run(stdinReader, printer)

		Eventually(buffer).Should(gbytes.Say("Welcome"))
		Eventually(buffer).Should(gbytes.Say("byeee"))
	})
})
