package compare

import (
	"github.com/trayo/go_mastermind/check"
	"github.com/trayo/go_mastermind/code_maker"
)

var (
	code  string
	guess string
)

type Compare struct {
	CorrectPositions string
	CorrectColors    string
	Code             string
}

func NewCompare(args ...string) *Compare {
	if len(args) > 1 {
		code = args[1]
	} else {
		code = code_maker.Generate()
	}

	return &Compare{
		CorrectPositions: "",
		CorrectColors:    "",
		Code:             code,
	}
}

func (self *Compare) Guess(guess string) {
	self.CorrectPositions = check.Positions(guess, self.Code)
	self.CorrectColors = check.Colors(guess, self.Code)
}
