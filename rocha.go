// Package rocha is a Hyperledger Fabric Chaincode Router with Middleware support
package rocha

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Router register routes to be invoked according to their method names
type Router struct {
	routes          map[string]Handler
	middlewares     []Middleware
	notFoundHandler Handler
}

// NewRouter creates a new Router
func NewRouter() *Router {
	return &Router{
		routes:          make(map[string]Handler),
		middlewares:     []Middleware{},
		notFoundHandler: NotFoundHandler,
	}
}

// Handle adds a new route to the router, overwriting a previous router by this name
func (r *Router) Handle(method string, handler Handler, middlewares ...Middleware) *Router {
	r.routes[method] = Chain(handler, middlewares...)
	return r
}

// NotFoundHandler sets a not found handler for this router
func (r *Router) NotFoundHandler(handler Handler) *Router {
	r.notFoundHandler = handler
	return r
}

// Use adds a new middleware
func (r *Router) Use(middlewares ...Middleware) *Router {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

// Invoke invokes the handler
func (r *Router) Invoke(stub shim.ChaincodeStubInterface, method string, args []string) pb.Response {
	// select which handler should invoke. If no handler is found, route
	// default notFoundHandler
	h, ok := r.routes[method]
	if !ok {
		h = r.notFoundHandler
	}

	// apply global middlewares to handler
	h = Chain(h, r.middlewares...)

	// create context, call handler and return the result back to the user
	return h(NewContext(stub, method, args))
}

// NotFoundHandler is a simple handler which returns a not found message with
// the method's name
func NotFoundHandler(c Context) pb.Response {
	return shim.Error(fmt.Sprintf("method '%s' not found", c.Method()))
}
