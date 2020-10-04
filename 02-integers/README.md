# 02 - Integers

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- when multiple arguments or return values use the same type, the type can
    use a shorthand where only a single type declarations follows a comma
    separated list of the names of the values:

    ```golang
    func myFunc(x, y int) {...}
    ```
- adding comments to functions in Go will add them to the generated
    documentation for the packages
    - `godoc` can be run in the directory where Go modules reside. The
        documentation can then be found at `http://localhost:6060` by default
    - documentation in `godoc` will include documentation for all packages in
        `$GOPATH`; you'll need to find your package in that documentation

### Tests

- examples for functions can be added in the associated `_test.go` file
  - to add an example:
    - prefix the function name with `Example`
    - add an example of the execution as the body of the function

        ```golang
        // somepackage_test.go
        // ...

        func ExampleMyFunc() {
          result := MyFunc(a ...interface{})
          // Output: expectation
        }
        ```
  - to add multiple examples, append `_[something]` to the function:

      ```golang
      func ExampleMyFunc_simple() {}
      func ExampleMyFunc_advanced() {}
      ```
  - by running `go test -v` one can see if the example is run or not
    - if an example does not contain a line having `// Output: [expectation]`
        then the example will not by executed. This helps ensure that examples
        provided with tests are valid as a codebase changes
    - try with examples that fail to see `go test`s output
