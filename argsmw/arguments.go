// Package argsmw is a chaincode argument parser and validator middleware
package argsmw

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/vtfr/rocha"
)

// Definition is a function which parses a chaincode argument at a given
// position and stores it's parsed value in a context key, if valid.
// Else, returns a error which will be returned to the Arguments middleware
// for further processing
type Definition func(c rocha.Context, arg string) error

// Arguments middleware is a argument parsing middleware which receives a set
// of argument definitions use them to parse the chaincode arguments. The
// parsed argument will be stored in their respective definition keys
func Arguments(defs ...Definition) rocha.Middleware {
	return func(next rocha.Handler) rocha.Handler {
		return func(c rocha.Context) pb.Response {
			// retrieve the invoke function arguments
			args := c.Args()

			// verify if they are the same length
			if len(args) != len(defs) {
				return shim.Error(
					fmt.Sprintf("Invalid number of arguments. Expected %d",
						len(defs)))
			}

			// for each argument, attempt to parse it's values
			for i, d := range defs {
				if err := d(c, args[i]); err != nil {
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
