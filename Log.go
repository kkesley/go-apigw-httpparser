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
	return Log(request, "REQUEST")
}

//LogResponse log a response
func LogResponse(code int, request interface{}) (err error) {
	return Log(request, fmt.Sprintf("RESPONSE-%d", code))
}

//Log the request
func Log(request interface{}, logType string) (err error) {
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
	fmt.Fprintf(os.Stderr, logType+": %s\n", string(b))
	return nil
}

func stripValues(t reflect.Type, v reflect.Value) (err error) {
	if v.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("log")
		if tag == "false" {
			value.Set(reflect.Zero(field.Type))
			continue
		}

		if value.Kind() == reflect.Struct {
			stripValues(value.Type(), value)
		} else if value.Kind() == reflect.Slice {
			for i := 0; i < value.Len(); i++ {
				stripValues(value.Index(i).Type(), value.Index(i).Elem())
			}
		} else if value.Kind() == reflect.Ptr && !value.IsNil() {
			stripValues(value.Elem().Type(), value.Elem())
		}

	}
	return nil
}
