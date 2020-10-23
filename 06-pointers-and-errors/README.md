# Structs, methods, and interfaces

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
    - [Pointers](#pointers)
    - [Error](#error)
  - [Tests](#tests)
- [Resources](#resources)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

#### Pointers

- go passes values as copies in function parameters by default
- the address of a value can be asserted by prepending the value with `&`
- modifying a copy of a value when that value is an instance property is going
    to do little
- Go uses `*` to make a value passed into a function point to the original
    value.

    ```golang
    func myFunc(v *MyValue) {}
    ```
- When defining methods, the receiver can either be a value receiver, or a
    pointer receiver:

    ```golang
    // pointter receiver
    func (t *MyType) MyMethod () {}

    // value receiver
    func (t *MyType) MyMethod () {}
    ```
- the receiver value of a method is equivalent to `this` in other languages
- as a convention, keep the types of receivers the same for a struct or
    interface. i.e., if one of the methods is a pointer receiver, then all of
    the methods should be pointer receivers, even if the pointer is not required
    internally
- inside a method that uses a pointer receiver, Go automatically dereferences
    the value, meaning that we don't need to  use `*` to reference the value -
    Go does this for us (and it's still valid to use the value with the
    pointer):

    ```golang
    func (t *MyType) MyMethod() {
        fmt.Println(t.myProperty)

        // is equivalent to
        fmt.Println((*t).myProperty)
    }
    ```

#### Error

- errcheck

### Tests


## Resources

- [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
