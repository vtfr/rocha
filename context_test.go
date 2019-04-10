package rocha_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
)

var _ = Describe("Context", func() {
	const METHOD = "method"
	ARGS := []string{"a", "b"}

	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, METHOD, ARGS)
	})

	It("Should return the stub related data", func() {
		Expect(context.Stub()).To(BeNil())
		Expect(context.Method()).To(Equal(METHOD))
		Expect(context.Args()).To(Equal(ARGS))
	})

	It("Should store and retrieve values", func() {
		const KEY = "key"
		const DATA = "VALUE"

		context.Set(KEY, DATA)
		Expect(context.Get(KEY)).To(Equal(DATA))
	})
})
