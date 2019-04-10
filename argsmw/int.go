package argsmw

import (
	"strconv"

	"github.com/vtfr/rocha"
)

// Int is a integer parses which stores a integer in the given context
// key. It also receives a base parameter which defines which base should
// be used when decoding the integer
func Int(key string, base int) Definition {
	return func(c rocha.Context, arg string) error {
		v, err := strconv.ParseInt(arg, base, 32)
		if err != nil {
			return err
		}

		c.Set(key, int(v))
		return nil
	}
}
