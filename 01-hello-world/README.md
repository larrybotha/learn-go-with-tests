# Hello World

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- a Go module may be initialised using `go mod init`, along with the repo the module
    belongs to, followed by the name of the module

    ```golang
    $ go mod init github.com/[username]/[repo|module]
    ```
- Go allows for _naked_ returns where the return values may be named as part of
    the functionn signature:

    ```golang
    func myFunc(param string) (value  string) {
      switch param {
      case "foo":
        value = "bar"
      default:
        value = "baz"
      }

      return
    }
    ```
- in a package, a definition starting with a capital letter is exported from the
    module. All definitions that start with a lowercase value are private to the
    package
- defining constants outside of functions seems to be a convention for
    optimising packages, so as to prevent re-initialisation of values every time
    a function is called
- functions can be defined withing functions using Go's assignment:

    ```golang
    func outerFunc() {
      innerFunc := func() {}
    }
    ```

### Tests

- test files should use the same name as that of the module they are testing,
    but with a `_test` suffix before the extension.
- tests import the `testing` package
- tests should start with `Test`

    ```golang
    func TestSomething(t *testing.T) {...}
    ```
- tests can be nested, similarly to `describe` blocks in Jest:

    ```golang
    func TestSomething(t *testing.T) {
      t.Run("message", func(t *testing.T) {
        ...
      })
    }
    ```
- tests can be failed using `t.Errorf("message", a ...interface{})`
- `t.Helper()` can be used inside test helper functions to indicate that the
    function is a helper function so that line and file information is not
    printed for executions inside that function
