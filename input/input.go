package input

import (
	"bufio"
	"strings"
)

func GetInput(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	return s
}

func WantsToQuit(s string) bool {
	if s == "q" || s == "quit" {
		return true
	}
	return false
}

func WantsInstructions(s string) bool {
	if s == "i" || s == "instructions" {
		return true
	}
	return false
}

func WantsToPlay(s string) bool {
	if s == "p" || s == "play" {
		return true
	}
	return false
}
