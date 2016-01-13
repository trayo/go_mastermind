package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/trayo/go_mastermind/print"
)

var (
	reader *bufio.Reader
	input  string
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	Run(stdin)
}

func Run(stdin *bufio.Reader) {
	reader = stdin

	print.WelcomeMessage()
	print.WhatsNext()
	input = getInput()

	for !wantsToQuit(input) {
		switch {
		case wantsInstructions(input):
			print.Instructions()
		case wantsToPlay(input):
			// game.Play()
			print.ThanksForPlaying()
		default:
			print.UnknownCommand()
		}

		print.WhatsNext()
		input = getInput()
	}

	print.QuitMessage()
}

func getInput() string {
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	return s
}

func sleep(s time.Duration) {
	time.Sleep(s * time.Millisecond)
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
