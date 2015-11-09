package code_maker

import (
	"math/rand"
	"time"
)

var (
	colors     = []string{"r", "b", "g", "y"}
	codeLength = 4
)

func Generate(params ...string) string {
type CodeMaker struct {
	code string
}

func NewCodeMaker() *CodeMaker {
	return &CodeMaker{}
}
	// if a specific code is passed in
	// use that as the code
	if len(params) >= 1 {
		return params[0]
	}

	rand.Seed(time.Now().UTC().UnixNano())
	code := ""

	for i := 0; i < easyModeCodeLength; i++ {
		code += colors[rand.Intn(len(colors))]
	}

	return code
}
