package rocha_test

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
)

var _ = Describe("Context", func() {
	STUB := &shim.MockStub{}
	METHOD := "method"
	ARGS := []string{"a", "b"}

	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(STUB, METHOD, ARGS)
	})

	It("Should return the chaincode related data", func() {
		Expect(context.Stub()).To(Equal(STUB))
		Expect(context.Method()).To(Equal(METHOD))
		Expect(context.Args()).To(Equal(ARGS))
	})

	It("Should return a value with Get", func() {
		context.Set("key", "data")
		value, exists := context.Get("key")
		Expect(value).To(Equal("data"))
		Expect(exists).To(BeTrue())

		value, exists = context.Get("invalid")
		Expect(value).To(BeNil())
		Expect(exists).To(BeFalse())
	})

	It("Should return a value with Value", func() {
		context.Set("key", "data")

		Expect(context.Value("key")).To(Equal("data"))

		Expect(context.Value("invalid")).To(BeNil())
	})

	It("Should return a value with String", func() {
		context.Set("key", "data")

		Expect(context.String("key")).To(Equal("data"))

		Expect(context.String("invalid")).To(BeEmpty())
	})

	It("Should return a value with Int", func() {
		context.Set("key", 10)

		Expect(context.Int("key")).To(Equal(int(10)))

		Expect(context.String("invalid")).To(BeZero())
	})
})
