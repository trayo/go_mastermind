package main

import (
	"bufio"
	"os"

	"github.com/trayo/go_mastermind/game"
	"github.com/trayo/go_mastermind/input"
	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	printer := print.NewPrinter(os.Stdout)
	gamer := game.NewGamer(stdin, printer)
	Run(gamer, stdin, printer)
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
		default:
			printer.UnknownMainCommand()
		}

		printer.WhatsNext()
		in = input.GetInput(stdin)
	}

	printer.QuitMessage()
}
