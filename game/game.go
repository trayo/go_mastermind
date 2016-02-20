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

type Gamer struct{}

func NewGamer() I.Gamer {
	return Gamer{}
}

func (g Gamer) Play(stdin *bufio.Reader, printer print.Printer) {
	printer.EnterAGuess()
	printer.ClearScreen()
}
