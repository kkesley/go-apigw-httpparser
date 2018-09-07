package httpparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
)

//LogRequest logs request to stderr
func LogRequest(request interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			if s, ok := r.(string); ok {
				panic(s)
			}
			err = r.(error)
		}
	}()
	value := reflect.ValueOf(request)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return errors.New("value must be pointer")
	}
	if err := stripValues(reflect.TypeOf(request).Elem(), value.Elem()); err != nil {
		return nil
	}
	b, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}
	fmt.Fprintf(os.Stderr, "REQUEST: %s\n", string(b))
	return nil
}

func stripValues(t reflect.Type, v reflect.Value) (err error) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Println(field.Type.Name())
		if value.Kind() == reflect.Struct {
			stripValues(value.Type(), value)
		} else if value.Kind() == reflect.Slice {
			for i := 0; i < value.Len(); i++ {
				stripValues(value.Index(i).Type(), value.Index(i))
			}
		}
		// Get the field tag value
		tag := field.Tag.Get("log")
		if tag != "false" {
			continue
		}
		value.Set(reflect.Zero(field.Type))
	}
	return nil
}