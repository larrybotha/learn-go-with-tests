# Dependency Injection

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

- types that implement specific interfaces can be used interchangeably in places
    where an implementation of an interface is expected. e.g. `fmt.Printf`
    is variadic, and accepts values of arbitrary types (due to its use of
    `a ...interface{}`), but it calls `fmt.Fprintf`. `fmt.Fprintf` expects a
    value of type `io.Writer` as its first argument. Instead of using
    `fmt.Printf` directly, we can call `fmt.Fprintf`, passing in whatever type
    that implements `io.Writer`:

    ```golang
    // This custom implementation is exactly what fmt.Printf does
    func MyCustomPrintf(format string, a ...interface{}) {
      // os.Stdout implements io.Writer's Write method
      return fmt.Fprintf(os.Stdout, format, a...)
    }
    ```

    and thus:

    ```golang
    func MyPrinter(w io.Writer, a ...interface{}) {
      return fmt.Fprintf(w, a...)
    }

    MyPrinter(os.Stdout, "hello")
    MyPrinter(*bytes.Buffer{}, "hello")
    ```
- by using dependency injection, one can move the responsilibity of an internal
    implementation of a function to the call-site. This can make it easier to
    test the outcome of testing a function. e.g. in React or Svelte, defining a
    default value for a property passed into a component, allowing a test to
    override the value if necessary, otherwise allowing the component to use the
    default value
- one can use `bytes.Buffer{}` to build strings

    ```golang
    b := bytes.Buffer{}

    // write to the buffer
    b.Write([]bytes("hello"))

    // or
    fmt.Fprintf(&b, ", world")

    // print to stdout
    b.WriteTo(os.Stdout)
    ```
- in the previous example, one needs to pass a pointer to `fmt.Fprintf` because
    of the way `bytes.Buffer` implements `io.Writer.Write`

    - `bytes.Buffer` implements `io.Writer.Write` using a pointer reference
    - if `b` were passed into `fmt.Fprintf` without the pointer, it would be
        passed in as a copy of the value of `b` - not the buffer to write to, which
        results in the following error:

        ```bash
        cannot use buffer (type bytes.Buffer) as type io.Writer in argument to [SomeFunc]:
        bytes.Buffer does not implement io.Writer (Write method has pointer receiver)
        ```
        - this error indicates that `bytes.Buffer.Write` has been implemented
        using a pointer receiver - it's implemented, but the way we've passed
        the buffer into the function is invalid
        - to fix this, when a method has been implemented with a pointer
        receiver, the value should be passed into call-sites as a pointer, using
        the & (ampersand) syntax
        - * operator is also termed as _the value at the address of_
- if an interface method is implemented with a pointer receiver, you have to pass
    a pointer to the value if you intend to use the interface
- if an interface method is implemented with a value receiver, you can pass
    either the value, or a pointer to the value - it doesn't matter

### Tests


## Resources

- [pointers in golang](https://www.geeksforgeeks.org/pointers-in-golang/)
- [answer in _Getting “bytes.Buffer does not implement io.Writer” error message_](https://stackoverflow.com/questions/23454940/getting-bytes-buffer-does-not-implement-io-writer-error-message#comment65376571_23454941)
