package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/trayo/go_mastermind/game"
	"github.com/trayo/go_mastermind/print"
)

var reader = *bufio.NewReader(os.Stdin)

func main() {
	clearScreen()
	print.WelcomeMessage()

	input := getInput()
	for !wantsToQuit(input) {
		switch {
		case wantsInstructions(input):
			clearScreen()
			print.Instructions()
		case wantsToPlay(input):
			clearScreen()
			print.GameStart()
			game.Play(&reader)
		default:
			print.UnknownCommand()
		}

		input = getInput()
	}

	print.QuitMessage()
	clearScreen()
	os.Exit(0)
}

func getInput() string {
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	return s
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
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
