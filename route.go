package rocha

import pb "github.com/hyperledger/fabric/protos/peer"

// Handler is a function called to process a chaincode request
type Handler func(c Context) pb.Response

// Middleware is function which wraps a Handler with some
// functionality, returning a new Handler with wrapped
// code
type Middleware func(next Handler) Handler

// Chain chains a series of middlewares to a new Handler. Middlewares
// are chained in the reverse order. If this function is called with
//
//    rocha.Chain(h, m4, m3, m2, m1) == m1(m2(m3(m4(h))))
//
func Chain(h Handler, middlewares ...Middleware) Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
