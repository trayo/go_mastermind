package game

import (
	"bufio"
	"os"

	"github.com/trayo/go_mastermind/check"
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
		in              string
		c               *check.Check
		guesses         int
		continuePlaying = true
	)

	if len(args) > 0 {
		c = check.NewCheck(args[0])
	} else {
		c = check.NewCheck()
	}

	g.printer.GameStart()

	for continuePlaying {
		g.printer.EnterAGuess()
		in = input.GetInput(g.stdin)
		switch {
		case input.WantsToQuit(in):
			continuePlaying = false
		case !input.Valid(in):
			g.printer.UnknownGameCommand()
		default:
			c.Guess(in)
			g.printer.CorrectColorsAndPositions(c.CorrectColors, c.CorrectPositions)
			guesses++
			if c.Won() {
				g.printer.YouWon(guesses)
				continuePlaying = false
			}
		}
	}

	g.printer.ThanksForPlaying()
}
