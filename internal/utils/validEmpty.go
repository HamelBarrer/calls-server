package utils

import (
	"fmt"
	"reflect"
)

func ValidEmpty[T comparable](m T) (string, bool) {
	k := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	for i := 0; i < k.NumField(); i++ {
		f := k.Field(i)
		va := v.Field(i)

		required := f.Tag.Get("required")
		if required == "true" && reflect.DeepEqual(va.Interface(), reflect.Zero(va.Type()).Interface()) {
			return fmt.Sprintf("%s is required", f.Name), true
		}
	}

	return "", false
}
