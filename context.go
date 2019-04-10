package rocha

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Context contains all necessary functions for interacting with the underlying
// blockchain and a simple key-value store for middleware/application usage
type Context interface {
	// Method returns this method's string name
	Method() string

	// Args returns the chaincode original arguments
	Args() []string

	// Stub returns the internal ChaincodeStub for this context
	Stub() shim.ChaincodeStubInterface

	// Set stores a value in the current context
	Set(key string, value interface{})

	// Get returns a value in the current context
	Get(key string) interface{}

	// GetString returns a string value in the current context
	GetString(key string) string

	// GetInt returns a integer value in the current context
	GetInt(key string) int
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

// Method returns this method's string name
func (c *context) Method() string { return c.method }

// Args returns the chaincode original arguments
func (c *context) Args() []string { return c.args }

// Stub returns the chaincode's shim.ChaincodeStubInterface
func (c *context) Stub() shim.ChaincodeStubInterface { return c.stub }

// Set stores a value in the current context
func (c *context) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get returns a value in the current context
func (c *context) Get(key string) interface{} {
	return c.data[key]
}

// GetString returns a string value in the current context
func (c *context) GetString(key string) string {
	return c.Get(key).(string)
}

// GetInt returns a integer value in the current context
func (c *context) GetInt(key string) int {
	return c.Get(key).(int)
}
