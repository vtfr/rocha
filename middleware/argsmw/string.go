package argsmw

import "github.com/vtfr/rocha"

// String
type String string

func (s String) Key() string {
	return string(s)
}

func (s String) Handle(c rocha.Context, arg string) error {
	c.Set(s.Key(), arg)
	return nil
}

// Compile-time verification if String implements Definition
var _ Definition = String("")
