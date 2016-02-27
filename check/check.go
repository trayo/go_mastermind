package check

import "github.com/trayo/go_mastermind/unique"

var (
	correctPositions int
	correctColors    int
)

func Positions(guess, code string) int {
	correctPositions = 0

	for i, letter := range guess {
		if uint8(letter) == code[i] {
			correctPositions++
		}
	}

	return correctPositions
}

func Colors(guess, code string) int {
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
