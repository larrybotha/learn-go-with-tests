# Maps

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)
- [Resources](#resources)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- maps are reference types
  - they are passed by reference, and the methods will mutate the map directly
- maps are defined by providing a type for the key and the value:

    ```golang
    myMap := map[string]int
    ```
- one should avoid creating maps with assigning any value. A map created without
    a default value references `nil` - attempting to assign values to nil will
    result in a runtime panic:

    ```golang
    // don't do this:
    myMap := map[string]string

    // instead, do this:
    myMap := map[string]string{}

    // or better, use make
    myMap := make(map[string]string)
    ```
- Go returns 2 values when retrieving a value at a key - the value, if it
    exists, and a boolean that can be used to check if the value exists:

    ```golang
    value, ok := myMap["foo"]
    ```
- deleting values from a map is idempotent; Go returns no value whether deleting
    with a key that exists or not
- there are some advantages to using constants for errors, such as errors being
    immutable. To create an error as a constant:
    - define a type alias for your errors on string
    - define an `Error()` method on that type - `Error()` is a method on the
        `error` interface

    ```golang
    type MyErr string

    const (
        ErrMyFirstError = MyErr("awesome error message")
        ErrMySecondError = MyErr("awesome other error message")
    )

    func (e MyErr) Error() string {
        return string(e)
    }
    ```

### Tests

- remember that a helper can take any arguments, so some of the repetitive work
    can be abstracted to the helper

## Resources

- [Constant errors](https://dave.cheney.net/2016/04/07/constant-errors)
