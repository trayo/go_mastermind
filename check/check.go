package check

import "fmt"

var (
	correctPositions int
	correctColors    int
)

func Positions(guess, code string) string {
	correctPositions = 0

	for i, letter := range guess {
		if uint8(letter) == code[i] {
			correctPositions++
		}
	}

	return fmt.Sprintf("%d", correctPositions)
}
