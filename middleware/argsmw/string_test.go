package argsmw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vtfr/rocha"
	. "github.com/vtfr/rocha/middleware/argsmw"
)

var _ = Describe("String", func() {
	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, "", []string{})
	})

	It("Should parse a string correctly", func() {
		err := String("key")(context, "value")
		Expect(err).To(BeNil())
		Expect(context.Value("key")).To(Equal("value"))
	})
})
