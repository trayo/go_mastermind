package compare

import (
	"github.com/trayo/go_mastermind/check"
	"github.com/trayo/go_mastermind/code_maker"
)

type Compare struct {
	CorrectPositions int
	CorrectColors    int
	Code             string
}

func NewCompare(args ...string) *Compare {
	var code string

	if len(args) >= 1 {
		code = args[0]
	} else {
		code = code_maker.Generate()
	}

	return &Compare{
		Code: code,
	}
}

func (c *Compare) Guess(guess string) {
	c.CorrectPositions = check.Positions(guess, c.Code)
	c.CorrectColors = check.Colors(guess, c.Code)
}
