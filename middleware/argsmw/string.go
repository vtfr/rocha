package argsmw

import "github.com/vtfr/rocha"

// String is a string parser which stores the string argument in a
// given context key
func String(key string) Definition {
	return func(c rocha.Context, arg string) error {
		// since it's already a string, no validation is needed
		c.Set(key, arg)
		return nil
	}
}
