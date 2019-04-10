package argsmw

import (
	"strconv"

	"github.com/vtfr/rocha"
)

// Int parses integers
type Int string

func (s Int) Key() string {
	return string(s)
}

func (s Int) Handle(c rocha.Context, arg string) error {
	v, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		return err
	}

	c.Set(s.Key(), v)
	return nil
}

// Compile-time verification if String implements Definition
var _ Definition = Int("")
