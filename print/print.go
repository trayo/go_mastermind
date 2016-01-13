package print

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
	clearScreen()
	fmt.Fprintln(p.writer, "Welcome to Mastermind!")
}

func (p Printer) QuitMessage() {
	fmt.Fprintln(p.writer, "byeeee")
}

func (p Printer) Instructions() {
	clearScreen()
	fmt.Fprintln(p.writer, "The game will generate a sequence with four elements made up of: ")
	fmt.Fprintln(p.writer, "(r)ed, (g)reen, (b)lue, and (y)ellow. It will look like 'rybg'.")
	fmt.Fprintln(p.writer, "Your goal is to guess the code in the fewest guesses possible.")
	fmt.Fprintln(p.writer, "\nGuesses must be four characters long and only contain 'r', 'g', 'b' and 'y'.")
	fmt.Fprintln(p.writer, "It doesn't matter if the guess is upper case or lower case.")
	fmt.Fprintln(p.writer, "The game will then tell you correct positions and correct colors.")
	fmt.Fprintln(p.writer, "Good luck!\n")
}

func (p Printer) UnknownCommand() {
	fmt.Fprintln(p.writer, "Unknown command! Please try again ...\n")
}

func (p Printer) WhatsNext() {
	fmt.Fprintln(p.writer, "Would you like to (p)lay read the (i)nstructions or (q)uit?")
	fmt.Fprintln(p.writer, "> ")
}

func (p Printer) GameStart() {
	clearScreen()
	fmt.Fprintln(p.writer, "I have generated a random code four characters in length using the letters:")
	fmt.Fprintln(p.writer, "'r', 'g', 'b' and 'y'")
	fmt.Fprintln(p.writer, "\nTry and guess the code by providing input like: 'rrrr'")
}

func (p Printer) EnterAGuess() {
	fmt.Fprintln(p.writer, "Enter your guess:")
	fmt.Fprintln(p.writer, "> ")
}

func (p Printer) ThanksForPlaying() {
	fmt.Fprintln(p.writer, "Thanks for playing!")
	fmt.Fprintln(p.writer, "Remember, winners don't do drugs.\n")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
