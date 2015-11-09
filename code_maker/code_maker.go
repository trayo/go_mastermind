package code_maker

import (
	"math/rand"
	"time"
)

var (
	colors     = []string{"r", "b", "g", "y"}
	codeLength = 4
)

type CodeMaker struct {
	code string
}

func NewCodeMaker() *CodeMaker {
	return &CodeMaker{}
}

func (self CodeMaker) Generate(params ...string) string {
	// if a specific code is passed in
	// use that as the code
	if len(params) >= 1 {
		return params[0]
	}

	rand.Seed(time.Now().UTC().UnixNano())
	self.code = ""

	for i := 0; i < codeLength; i++ {
		self.code += colors[rand.Intn(len(colors))]
	}

	return self.code
}
