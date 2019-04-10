package argsmw_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArgsmw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Argsmw Suite")
}
