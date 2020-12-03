# Reflection

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- one can use `interface{}` in Go much in the same way as one can use `any` in
    Typescript
    - consequently, it should be avoided
- to inspect / evaluate values that are passed into functions that use `interface{}` as a
    parameter, one can use `reflect`
- `reflect.ValueOf(x)` returns a `reflect.Value` that describes the value
- `reflect.Value.Kind()` returns the type of `reflect.Value`. This could be one
    of:
    - `reflect.String`
    - `reflect.Slice`
    - `reflect.Array`
    - `reflect.Map`
    - `reflect.Chan`
    - `reflect.Func`
- for a `reflect.Value` that is an array or slice:
  - one can retrieve the number of elements in the slice / array using
    `reflect.Value.Len()`
  - one can retrieve individual items using `reflect.Value.Index(index int)`
- for `reflect.Value` that is a map:
    - one can retrieve all the keys in the map using `reflect.Value.MapKeys()`
    - as we are evaluating a map, one should not rely on the order of keys
    - to get a value from a key, one can use `reflect.Value.MapIndex(key
        reflect.Value)`
- for `reflect.Value` that is a chan:
    - one can retrieve individual values using `reflect.Value.Recv`. As with
        pulling values from a channel, to get all values in the channel,
        `reflect.Value.Recv()` must be called until it no longer contains any
        values:

        ```golang
        for x, ok := val.Recv(); ok; v, ok = val.Recv() {
            // do something with x
        }
        ```
- for `reflect.Value` that is a function:
    - one can retrieve `[]reflect.Value` for the return value of the function
        using `reflect.Value.Call(nil)` (given a function that expects no arguments)
- to get the underlying actual value from `reflect.Value`, use
    `value.Interface()`

### Tests

- remember that helper functions in Go tests should be passed the test
    instance, and the helper should contain a call to `t.Helper()`

