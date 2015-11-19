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

func NewCompare(params ...string) *Compare {
	guess := params[0]
	code := params[1]

	if params[1] == nil {
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
