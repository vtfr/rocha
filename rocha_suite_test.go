package rocha_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRocha(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rocha Suite")
}
