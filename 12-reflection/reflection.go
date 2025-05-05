package reflection

import (
	"reflect"
)

func walk(x any, fn func(y string)) {
	value := getValue(x)
	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())

	case reflect.Struct:
		for i := range value.NumField() {
			walkValue(value.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := range value.Len() {
			walkValue(value.Index(i))
		}

	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}

	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		result := value.Call(nil)

		for _, x := range result {
			walkValue(x)
		}
	}
}

func getValue(x any) reflect.Value {
	value := reflect.ValueOf(x) // get a reflect.Value from the variable

	if value.Kind() == reflect.Pointer {
		// dereference the value
		value = value.Elem()
	}

	return value
}
