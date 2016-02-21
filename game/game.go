package game

import (
	"bufio"
	"os"

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
	g.printer.GameStart()
	g.printer.EnterAGuess()
	g.printer.ClearScreen()
}
