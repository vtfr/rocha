package argsmw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
	"github.com/vtfr/rocha/argsmw"
)

var _ = Describe("Int", func() {
	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, "", []string{})
	})

	It("Should parse integers correctly", func() {
		var err error

		// parses a base 10 integer and stores at key `value10`
		err = argsmw.Int("value10", 10)(context, "1337")
		Expect(err).ToNot(HaveOccurred())
		Expect(context.Int("value10")).To(Equal(int(1337)))

		// parses a base 16 integer and stores at key `value16`
		err = argsmw.Int("value16", 16)(context, "539")
		Expect(err).ToNot(HaveOccurred())
		Expect(context.Int("value10")).To(Equal(int(0x539)))
	})

	It("Should return an error if can't parse", func() {
		err := argsmw.Int("key", 10)(context, "15abc")

		Expect(err).To(HaveOccurred())
		Expect(context.Value("key")).To(BeNil())
	})
})
