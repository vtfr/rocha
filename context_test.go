package rocha_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
)

var _ = Describe("Context", func() {
	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, "method", []string{})
	})

	It("Should store and retrieve values", func() {
		PAYLOAD := interface{}("VALUE")
		KEY := "key"

		context.Set(KEY, PAYLOAD)
		Expect(context.Get(KEY)).To(Equal(PAYLOAD))
	})
})
