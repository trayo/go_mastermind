package compare

import (
	"github.com/trayo/go_mastermind/check"
	"github.com/trayo/go_mastermind/code_maker"
)

type Compare struct {
	CorrectPositions string
	CorrectColors    string
	code             string
}

func NewCompare(args ...string) *Compare {
	var (
		code  string
		guess = args[0]
	)

	if len(args) > 1 {
		code = args[1]
	} else {
		code = code_maker.Generate()
	}

	return &Compare{
		CorrectPositions: check.Positions(guess, code),
		CorrectColors:    check.Colors(guess, code),
		code:             code,
	}
}

func (self *Compare) Guess(guess string) {
	self.CorrectPositions = check.Positions(guess, self.code)
	self.CorrectColors = check.Colors(guess, self.code)
}
