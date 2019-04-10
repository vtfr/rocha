package rocha

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Context contains all necessary functions for interacting with the underlying
// blockchain infrastructure and a simple key-value store for
// middleware/application internal usage
type Context interface {
	// Method returns the method's name
	Method() (method string)

	// Args returns the sent arguments
	Args() (args []string)

	// Stub returns the internal ChaincodeStub for this context
	Stub() shim.ChaincodeStubInterface

	// Set stores a value in the current context
	Set(key string, value interface{})

	// Get returns a value in the current context
	Get(key string) (value interface{}, exists bool)

	// Value returns a value stored in the context like Get, but does no
	// checking if the value actually exists
	Value(key string) (value interface{})

	// String returns a value stored in the context casted to a string.
	// if no value is found, a empty string is returned
	String(key string) (value string)

	// Int returns a value stored in the context casted to a int. If no
	// value is found, a zero int is returned
	Int(key string) (value int)
}

// context is the default rocha.Context implementation
type context struct {
	stub   shim.ChaincodeStubInterface
	method string
	args   []string
	data   map[string]interface{}
}

// NewContext creates a new context with the middlewares and handlers to be called
func NewContext(stub shim.ChaincodeStubInterface, method string, args []string) Context {
	return &context{
		stub:   stub,
		method: method,
		args:   args,
		data:   make(map[string]interface{}),
	}
}

func (c *context) Method() string { return c.method }

func (c *context) Args() []string { return c.args }

func (c *context) Stub() shim.ChaincodeStubInterface { return c.stub }

func (c *context) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.data[key]
	return
}

func (c *context) Value(key string) interface{} {
	return c.data[key]
}

func (c *context) String(key string) string {
	value, _ := c.data[key].(string)
	return value
}

func (c *context) Int(key string) int {
	value, _ := c.data[key].(int)
	return value
}
