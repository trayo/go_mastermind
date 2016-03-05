package check

import (
	"github.com/trayo/go_mastermind/code_maker"
	"github.com/trayo/go_mastermind/unique"
)

var (
	correctPositions int
	correctColors    int
)

type Check struct {
	CorrectPositions int
	CorrectColors    int
	code             string
}

func NewCheck(args ...string) *Check {
	var code string

	if len(args) >= 1 {
		code = args[0]
	} else {
		code = code_maker.Generate()
	}

	return &Check{
		code: code,
	}
}

func (c *Check) Guess(guess string) {
	c.CorrectPositions = findPositions(guess, c.code)
	c.CorrectColors = findColors(guess, c.code)
}

func (c *Check) Won() bool {
	return c.CorrectPositions == 4 && c.CorrectColors == 4
}

func findPositions(guess, code string) int {
	correctPositions = 0

	for i, letter := range guess {
		if uint8(letter) == code[i] {
			correctPositions++
		}
	}

	return correctPositions
}

func findColors(guess, code string) int {
	correctColors = 0

	uniqueGuess := unique.RemoveDuplicateLetters(guess)

	for _, guessLetter := range uniqueGuess {
		for _, codeLetter := range code {
			if guessLetter == codeLetter {
				correctColors += 1
			}
		}
	}

	return correctColors
}
