package utils

import (
	"fmt"
	"reflect"
)

func SetField[T any, V any](obj *T, field string, value V) {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(field)

	if !structFieldValue.IsValid() {
		fmt.Println("No such field:", field)
		return
	}
	if !structFieldValue.CanSet() {
		fmt.Println("Cannot set value for field:", field)
		return
	}
	val := reflect.ValueOf(value)
	if structFieldValue.Type() != val.Type() {
		fmt.Println("Provided value type doesn't match field type")
		return
	}
	structFieldValue.Set(val)
}
