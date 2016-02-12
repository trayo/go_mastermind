package main

import (
	"bufio"
	"strings"

	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"
)

var (
	reader *bufio.Reader
	input  string
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	printer := print.NewPrinter(os.Stdout)
	Run(stdin, printer)
}

func Run(gamer I.Gamer, stdin *bufio.Reader, printer print.Printer) {
	reader = stdin

	printer.ClearScreen()
	printer.WelcomeMessage()
	printer.WhatsNext()
	input = getInput()

	for !wantsToQuit(input) {
		switch {
		case wantsInstructions(input):
			printer.ClearScreen()
			printer.Instructions()
		case wantsToPlay(input):
			printer.ClearScreen()
			gamer.Play()
			printer.ThanksForPlaying()
		default:
			printer.UnknownCommand()
		}

		printer.WhatsNext()
		input = getInput()
	}

	printer.QuitMessage()
}

func getInput() string {
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	return s
}

func wantsToQuit(s string) bool {
	if s == "q" || s == "quit" {
		return true
	}
	return false
}

func wantsInstructions(s string) bool {
	if s == "i" || s == "instructions" {
		return true
	}
	return false
}

func wantsToPlay(s string) bool {
	if s == "p" || s == "play" {
		return true
	}
	return false
}
