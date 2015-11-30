package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoMastermind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mastermind Suite")
}
