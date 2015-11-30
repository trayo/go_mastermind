package main_test

import (
	"bufio"
	"bytes"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/trayo/go_mastermind"
)

var _ = Describe("Main CLI interaction", func() {

	var (
		stdin       bytes.Buffer
		stdinReader bufio.Reader

		stdout       bytes.Buffer
		stdoutReader bufio.Reader
	)

	BeforeEach(func() {
		stdin = bytes.Buffer{}
		stdinReader = *bufio.NewReader(&stdin)

		stdout = bytes.Buffer{}
		stdoutReader = *bufio.NewReader(&stdout)
	})

	It("can quit the game", func() {
		stdin.WriteString("q")

		main.Run(&stdinReader, &stdoutReader)

		stdoutReader.ReadString('\n')

		fmt.Println("reader: ", stdout)
		Expect(stdout).To(Equal("instructions"))
	})
})
