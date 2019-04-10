package argsmw_test

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vtfr/rocha"
	"github.com/vtfr/rocha/middleware/argsmw"
)

var _ = Describe("Argument", func() {
	PAYLOAD := shim.Success([]byte("OK"))
	handler := rocha.Handler(func(c rocha.Context) pb.Response {
		return PAYLOAD
	})

	It("Should parse arguments correctly", func() {
		context := rocha.NewContext(nil, "", []string{"string1", "string2"})

		mw := argsmw.Arguments(argsmw.String("key1"), argsmw.String("key2"))
		ret := mw(handler)(context)

		Expect(ret).To(Equal(PAYLOAD))
		Expect(context.Value("key1")).To(Equal("string1"))
		Expect(context.Value("key2")).To(Equal("string2"))
	})

	It("Should return a error message if a invalid number of argument is sent", func() {
		context := rocha.NewContext(nil, "", []string{"string1"})

		mw := argsmw.Arguments(argsmw.String("key1"), argsmw.String("key2"))
		ret := mw(handler)(context)

		Expect(int(ret.Status)).To(Equal(shim.ERROR))
		Expect(ret.Message).To(ContainSubstring("Invalid number of arguments"))
	})

	It("Should return a error message can't parse an argument", func() {
		context := rocha.NewContext(nil, "", []string{"10a"})

		mw := argsmw.Arguments(argsmw.Int("key", 10))
		ret := mw(handler)(context)

		Expect(int(ret.Status)).To(Equal(shim.ERROR))
		Expect(ret.Message).To(ContainSubstring("Invalid argument at position"))
	})
})
