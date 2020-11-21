package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	numValues := 0
	var getField func(int) reflect.Value

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
		numValues = val.Len()
		getField = val.Index

	case reflect.Array:
		numValues = val.Len()
		getField = val.Index

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}

	case reflect.Func:
		result := val.Call(nil)

		for _, x := range result {
			walk(x.Interface(), fn)
		}
	}

	for i := 0; i < numValues; i++ {
		field := getField(i)
		walk(field.Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
