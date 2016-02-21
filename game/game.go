package game

import (
	"bufio"
	"os"

	"github.com/trayo/go_mastermind/input"
	I "github.com/trayo/go_mastermind/interfaces"
	"github.com/trayo/go_mastermind/print"
)

var (
	reader = *bufio.NewReader(os.Stdin)
)

type Gamer struct {
	stdin   *bufio.Reader
	printer print.Printer
}

func NewGamer(stdin *bufio.Reader, printer print.Printer) I.Gamer {
	return Gamer{
		stdin:   stdin,
		printer: printer,
	}
}

func (g Gamer) Play(args ...string) {
	var (
		in string
	)
	g.printer.GameStart()
	g.printer.EnterAGuess()
	in = input.GetInput(g.stdin)

	for !input.WantsToQuit(in) {

		g.printer.EnterAGuess()
		in = input.GetInput(g.stdin)
	}

	g.printer.ThanksForPlaying()
}
