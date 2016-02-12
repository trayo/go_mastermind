package code_maker

import (
	"math/rand"
	"time"
)

const (
	CODELENGTH = 4
)

var COLORS = []string{"r", "b", "g", "y"}

func Generate(params ...string) string {
	// if a specific code is passed in
	// use that as the code
	if len(params) >= 1 {
		return params[0]
	}

	rand.Seed(time.Now().UTC().UnixNano())
	code := ""

	for i := 0; i < CODELENGTH; i++ {
		code += COLORS[rand.Intn(len(COLORS))]
	}

	return code
}
