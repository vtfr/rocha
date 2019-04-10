package argsmw

import (
	"encoding/json"
	"reflect"

	"github.com/vtfr/rocha"
)

// JSON is a JSON parser which receives a base structure and parses a
// argument according to it's format.
//
//     r.Handler(handler, argsmw.Arguments(argsmw.JSON("value", &Request{})))
//
// And, in the handler:
//
//     func handler(c rocha.Context) pb.Response {
//         request := c.Value("value").(*Request)
//         ...
//     }
func JSON(key string, model interface{}) Definition {
	// verifies if is a pointer
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		panic("argwm.JSON: value must be a pointer to a struct")
	}

	// stores the data type
	modelType := reflect.ValueOf(model).Elem().Type()

	return func(c rocha.Context, arg string) error {
		// clones the data by it's type
		value := reflect.New(modelType).Interface()

		// parses JSON
		if err := json.Unmarshal([]byte(arg), &value); err != nil {
			return err
		}

		c.Set(key, value)
		return nil
	}
}
