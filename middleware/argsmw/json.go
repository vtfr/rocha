// Package argsmw
package argsmw

import (
	"encoding/json"
	"reflect"

	"github.com/vtfr/rocha"
)

type jsondef struct {
	key   string
	model reflect.Type
}

// JSON parses a JSON encoded argument to a specific class and stores the
// parsed value at the given key
func JSON(key string, value interface{}) Definition {
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		panic("argwm.JSON: value must be a pointer to a struct")
	}

	return &jsondef{key, reflect.ValueOf(value).Elem().Type()}
}

// Key returns in which key the parsed JSON element will be stored
func (j *jsondef) Key() string { return j.key }

// Handle parses the JSON
func (j *jsondef) Handle(c rocha.Context, arg string) error {
	value := reflect.New(j.model).Interface()

	// attempt to unmarshal, else return the error
	if err := json.Unmarshal([]byte(arg), &value); err != nil {
		// TODO(vtfr): wrap error
		return err
	}

	c.Set(j.key, value)
	return nil
}
