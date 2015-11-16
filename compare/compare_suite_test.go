package compare_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompare(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compare Suite")
}
