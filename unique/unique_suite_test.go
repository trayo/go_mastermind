package unique_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnique(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unique Suite")
}
