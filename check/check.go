package check

import (
	"github.com/trayo/go_mastermind/code_maker"
	"github.com/trayo/go_mastermind/unique"
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
	return c.CorrectPositions == len(c.code)
}

func findPositions(guess, code string) int {
	correctPositions := 0

	for i, letter := range guess {
		if uint8(letter) == code[i] {
			correctPositions++
		}
	}

	return correctPositions
}

func findColors(guess, code string) int {
	uniqueGuess := unique.RemoveDuplicateLetters(guess)
	uniqueCode := unique.RemoveDuplicateLetters(code)

	var correctColors int
	for _, guessLetter := range uniqueGuess {
		for _, codeLetter := range uniqueCode {
			if guessLetter == codeLetter {
				correctColors++
			}
		}
	}

	return correctColors
}
