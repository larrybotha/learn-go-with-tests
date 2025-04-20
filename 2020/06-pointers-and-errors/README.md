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

- a type alias is defined in the following way:

    ```golang
    type [TypeAlias] [originalType]
    ```
- the `Stringer` interface has a `String` method which allows one to customise
    how values are printed. e.g.

    ```golang
    // create a type alias
    type MyMoneyType int

    func (m MyMoneyType) String() {
        fmt.Printf("%d ChaChing!", m)
    }
    ```

    - values printed with formatting need to be printed with the `%s` flag to
        indicate the value should be printed as a string


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
    // pointer receiver
    func (t *MyType) MyMethod () {}

    // value receiver
    func (t MyType) MyMethod () {}
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
- using `var` to define a top-level value in a package will make that variable
    available outside of the package

#### Error

- error messages are represented by the builtin interface `error`
- `error.Error()` can be called to get the message in an error
- functions often return an error, which should be evaluated to determine if an
    error has been thrown, otherwise the app crash due to runtime errors
- [errcheck](https://github.com/kisielk/errcheck) can be used to determine if
    one's code has not handled any returned errors

    ```bash
    # run in current directory where a function is not handdling a returned
    # error:
    $ errcheck
    wallet_test.go:35:18: wallet.Withdraw(Bitcoin(5))
    ```
- errors can be created using `errors.New("error message")`:

    ```golang
    // make InvalidPassword available globally
    var InvalidPassword = errors.New("Invalid password")
    ```

    - this type of error is called a sentinel - see [this article](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
    - a sentinel is a named error that is exported. The problem with exporting
        an error like this is it ties other packages to your package - if the
        error changes or is removed, packages relying on it will break. Avoid
        exporting sentinels
- errors are not thrown in Go - one returns errors from inside functions, and
    that error should be evaluated at the call site

### Tests

- helpers can be defined outside of the actual tests. By placing helpers after
    our tests, we make it easier for users who are evaluating our packages to
    read our tests instead of wading through the noise of helpers

## Resources

- [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
