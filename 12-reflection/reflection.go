package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}

	case reflect.Func:
		result := val.Call(nil)

		for _, x := range result {
			walkValue(x)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
