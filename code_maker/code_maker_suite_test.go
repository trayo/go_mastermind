package code_maker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCodeMaker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeMaker Suite")
}
