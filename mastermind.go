package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/trayo/go_mastermind/game"
	"github.com/trayo/go_mastermind/print"
)

var (
	reader = *bufio.NewReader(os.Stdin)
	input  string
)

func main() {
	print.WelcomeMessage()

	for !wantsToQuit(input) {
		print.WhatsNext()
		input = getInput()
		switch {
		case wantsInstructions(input):
			print.Instructions()
		case wantsToPlay(input):
			game.Play()
			print.ThanksForPlaying()
		default:
			print.UnknownCommand()
		}

	}

	print.QuitMessage()
	os.Exit(0)
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
