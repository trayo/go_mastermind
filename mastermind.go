package main

import (
	"bufio"

	"github.com/trayo/go_mastermind/input"
	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	printer := print.NewPrinter(os.Stdout)
	Run(stdin, printer)
}

func Run(gamer I.Gamer, stdin *bufio.Reader, printer print.Printer) {
	var in string

	printer.ClearScreen()
	printer.WelcomeMessage()
	printer.WhatsNext()
	in = input.GetInput(stdin)

	for !input.WantsToQuit(in) {
		switch {
		case input.WantsInstructions(in):
			printer.ClearScreen()
			printer.Instructions()
		case input.WantsToPlay(in):
			printer.ClearScreen()
			gamer.Play()
			printer.ThanksForPlaying()
		default:
			printer.UnknownCommand()
		}

		printer.WhatsNext()
		in = input.GetInput(stdin)
	}

	printer.QuitMessage()
}
