// Package argsmw
package argsmw

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/vtfr/rocha"
)

// Definition is a argument definition containg a Key in which the argument
// will be stored and a Handle function for parsing the argument
type Definition interface {
	// Key returns the Context.Key in which this definition will store the
	// parsed value
	Key() string

	// Handle handles the argument
	Handle(rocha.Context, string) error
}

// ErrInvalidArgumentCount is returned when a invalid argument is sent to a
// Method
var ErrInvalidArgumentCount = errors.New("invalid argument count")

// Arguments middleware is a argument parsing middleware which receives a set
// of argument definitions and parse them, saving their parsed values to their
// respective keys
func Arguments(defs ...Definition) rocha.Middleware {
	return func(next rocha.Handler) rocha.Handler {
		return func(c rocha.Context) pb.Response {
			// retrieve the invoke function arguments
			args := c.Args()

			// verify if they are the same length
			if len(args) != len(defs) {
				return shim.Error(
					fmt.Sprintf("Invalid argument count. Expected %d",
						len(defs)))
			}

			// for each argument, attempt to parse it's values
			for i, def := range defs {
				if err := def.Handle(c, args[i]); err != nil {
					return shim.Error(
						fmt.Sprintf("Invalid argument at position '%d': %s",
							i, err.Error()))
				}
			}

			// call the next handler and return it's response
			return next(c)
		}
	}
}
