package rocha_test

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vtfr/rocha"
)

var _ = Describe("Chain", func() {
	It("Should chain requests in the correct order", func() {
		PAYLOAD := shim.Success([]byte("OK"))
		EXPECTED := []string{"M1", "M2", "M3", "H"}
		received := []string{}

		// default payload which simply returns the payload
		handler := func(_ rocha.Context) pb.Response {
			received = append(received, "H")
			return PAYLOAD
		}

		// middleware factory which generates a simple tagger middleware
		tagMiddleware := func(tag string) rocha.Middleware {
			return func(next rocha.Handler) rocha.Handler {
				return func(c rocha.Context) pb.Response {
					received = append(received, tag)
					return next(c)
				}
			}
		}

		m1 := tagMiddleware("M1")
		m2 := tagMiddleware("M2")
		m3 := tagMiddleware("M3")

		h := rocha.Chain(handler, m3, m2, m1)
		res := h(rocha.NewContext(nil, "", []string{}))
		Expect(res).To(Equal(PAYLOAD))
		Expect(received).To(Equal(EXPECTED))
	})
})
