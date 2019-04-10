package argsmw_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
	"github.com/vtfr/rocha/argsmw"
)

type testJSONContent struct {
	Parameter1 string `json:"param1"`
	Parameter2 int    `json:"param2"`
}

var _ = Describe("JSON", func() {
	var context rocha.Context

	BeforeEach(func() {
		context = rocha.NewContext(nil, "", []string{})
	})

	It("Should parse JSON correctly", func() {
		JSON_CONTENT := `{"param1":"test","param2":10}`

		// test parsing the JSON content
		err := argsmw.JSON("key", &testJSONContent{})(context, JSON_CONTENT)
		Expect(err).ToNot(HaveOccurred())

		// test accesing the JSON element
		data, ok := context.Value("key").(*testJSONContent)
		Expect(ok).To(BeTrue())
		Expect(data.Parameter1).To(Equal("test"))
		Expect(data.Parameter2).To(Equal(10))
	})

	It("Should return error at invalid JSON", func() {
		JSON_CONTENT := `{"param5":"test",,,,}`

		// test parsing the JSON content
		err := argsmw.JSON("key", &testJSONContent{})(context, JSON_CONTENT)
		Expect(err).To(HaveOccurred())

		// test if key is unchanged
		Expect(context.Value("key")).To(BeNil())
	})

	It("Should panic if called with invalid model struct", func() {
		defer func() {
			msg := recover()
			Expect(msg).To(ContainSubstring("value must be a pointer"))
		}()

		argsmw.JSON("key", testJSONContent{})
	})
})
