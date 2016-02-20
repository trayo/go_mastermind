package print

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

type Printer struct {
	writer io.Writer
}

func NewPrinter(writer io.Writer) Printer {
	return Printer{
		writer: writer,
	}
}

func (p Printer) WelcomeMessage() {
	fmt.Fprintln(p.writer, "Welcome to Mastermind!")
}

func (p Printer) QuitMessage() {
	fmt.Fprintln(p.writer, "byeeee")
}

func (p Printer) Instructions() {
	fmt.Fprintln(p.writer, "The game will generate a sequence with four elements made up of: ")
	fmt.Fprintf(p.writer, "(%s)ed, (%s)reen, (%s)lue, and (%s)ellow. It will look like '%s'.\n", colorize("r"), colorize("g"), colorize("b"), colorize("y"), colorize("rrrr"))
	fmt.Fprintln(p.writer, "Your goal is to guess the code in the fewest guesses possible.")
	fmt.Fprintf(p.writer, "\nGuesses must be four characters long and only contain %s", colorize("'r', 'g', 'b' and 'y'.\n"))
	fmt.Fprintln(p.writer, "It doesn't matter if the guess is upper case or lower case.")
	fmt.Fprintln(p.writer, "The game will then tell you correct positions and correct colors.")
	fmt.Fprintln(p.writer, "Good luck!\n")
}

func (p Printer) UnknownCommand() {
	fmt.Fprintln(p.writer, "Unknown command! Please try again ...\n")
}

func (p Printer) WhatsNext() {
	fmt.Fprintln(p.writer, "Would you like to (p)lay read the (i)nstructions or (q)uit?")
	fmt.Fprint(p.writer, "> ")
}

func (p Printer) GameStart() {
	fmt.Fprintln(p.writer, "I have generated a random code four characters in length using the letters:")
	fmt.Fprintf(p.writer, colorize("'r', 'g', 'b' and 'y'"))
	fmt.Fprintln(p.writer, "\nTry and guess the code by providing input like: %s", colorize("'rrrr'"))
}

func (p Printer) EnterAGuess() {
	fmt.Fprintln(p.writer, "Enter a guess:")
	fmt.Fprint(p.writer, "> ")
}

func (p Printer) ThanksForPlaying() {
	fmt.Fprintln(p.writer, "Thanks for playing!")
	fmt.Fprintln(p.writer, "Remember, winners don't do drugs.\n")
}

func (p Printer) ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = p.writer
	c.Run()
}

func colorize(guess string) string {
	var result string

	for _, s := range strings.Split(guess, "") {
		switch s {
		default:
			result += s
		case "r":
			result += color.RedString(s)
		case "g":
			result += color.GreenString(s)
		case "y":
			result += color.YellowString(s)
		case "b":
			result += color.BlueString(s)
		}
	}

	return result
}
