package rocha_test

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vtfr/rocha"
)

var _ = Describe("Router", func() {
	It("Should route correctly", func() {
		// respond is a handler factory which generate handlers that
		// return a fixed response
		respond := func(value pb.Response) rocha.Handler {
			return func(_ rocha.Context) pb.Response {
				return value
			}
		}

		PAYLOAD1 := shim.Success([]byte("First Payload"))
		PAYLOAD2 := shim.Success([]byte("Second Payload"))
		PAYLOAD3 := shim.Success([]byte("Third Payload"))
		PAYLOAD4 := shim.Success([]byte("Not found"))

		h1 := respond(PAYLOAD1)
		h2 := respond(PAYLOAD2)
		h3 := respond(PAYLOAD3)
		h4 := respond(PAYLOAD4)

		router := rocha.NewRouter().
			Handle("Method1", h1).
			Handle("Method2", h2).
			Handle("Method3", h3).
			NotFoundHandler(h4)

		Expect(router.Invoke(nil, "Method1", []string{})).To(Equal(PAYLOAD1))
		Expect(router.Invoke(nil, "Method2", []string{})).To(Equal(PAYLOAD2))
		Expect(router.Invoke(nil, "Method3", []string{})).To(Equal(PAYLOAD3))
		Expect(router.Invoke(nil, "Method4", []string{})).To(Equal(PAYLOAD4))
	})

	It("Should handle middlewares", func() {
		PAYLOAD := shim.Success([]byte("OK"))

		middlewareCalled := 0

		middleware := func(next rocha.Handler) rocha.Handler {
			return func(c rocha.Context) pb.Response {
				middlewareCalled++
				return next(c)
			}
		}

		handler := func(c rocha.Context) pb.Response {
			return PAYLOAD
		}

		response := rocha.NewRouter().
			Use(middleware).
			Handle("Method", handler).
			Invoke(nil, "Method", []string{})

		Expect(response).To(Equal(PAYLOAD))
		Expect(middlewareCalled).To(Equal(1))
	})
})
