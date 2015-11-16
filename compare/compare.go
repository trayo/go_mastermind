package compare

import "github.com/trayo/go_mastermind/check"

type Compare struct {
	CorrectPositions string
	CorrectColors    string
	code             string
}

func NewCompare(guess, code string) *Compare {
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
