package struc

import (
	"fmt"
	"reflect"
)

func InsepectStruct(v any) string {
	valStruct := reflect.ValueOf(v)
	typeStruct := reflect.TypeOf(v)
	var description string

	if typeStruct.Kind() == reflect.Ptr {
		valStruct = valStruct.Elem()
		typeStruct = typeStruct.Elem()
	}
	for i := 0; i < typeStruct.NumField(); i++ {
		value := valStruct.Field(i)
		field := typeStruct.Field(i)
		description = description + fmt.Sprintf("Field: %s, Type: %s, Value: %v\n", field.Name, field.Type, value.Interface())
	}
	return description
}
