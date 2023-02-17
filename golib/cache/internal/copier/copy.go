package copier

import (
	"fmt"
	"reflect"

	"sadlil.com/samples/golib/cache"
)

// Copy copies the value from the source variable to the destination variable,
// preserving the underlying type and structure. Both source and destination
// must be pointers to structs with identical type and field names. If any field
// in the source struct is unexported (i.e., its name starts with a lowercase
// letter), it will be ignored during the copy. The function returns an error
// if the source or destination is not a pointer to a struct, or if the structs
// are not identical.
func Copy(src, dest interface{}) error {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest)
	if srcValue.Kind() != destValue.Kind() {
		return fmt.Errorf("type mismatched: %w", cache.ErrInvalidObject)
	}

	if srcValue.Kind() != reflect.Ptr {
		return cache.ErrInvalidObject
	}

	srcValue = srcValue.Elem()
	destValue = destValue.Elem()

	srcType := srcValue.Type()
	destType := destValue.Type()
	if srcType != destType {
		return fmt.Errorf("type mismatched: %w", cache.ErrInvalidObject)
	}

	for i := 0; i < srcValue.NumField(); i++ {
		srcFieldValue := srcValue.Field(i)
		destFieldValue := destValue.Field(i)

		if !destFieldValue.CanSet() {
			continue
		}

		if srcFieldValue.Type() != destFieldValue.Type() {
			continue
		}
		destFieldValue.Set(srcFieldValue)
	}
	return nil
}
