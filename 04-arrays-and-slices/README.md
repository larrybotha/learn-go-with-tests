# Arrays and slices

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- an array in Go has a fixed length:

    ```golang
    xs = [5]int{1,2,3,4,5}
    // or
    xs = [...]int{1,2,3,4,5}
    ```
- a slice in Go has a variable length, and can be defined by omitting the array
    parameter when defining the value:

    ```golang
    xs = []int{1,2,3,4,5}
    ```
- `for` loops allow for the use of a `range` statement that will iterate over an
    array or slice, returning an index and value for each iteration:

    ```golang
    xs = []int{1,2,3,4,5}

    for i, number := range {
    fmt.Printf("%d => %d", i, number)
    }
    ```

### Tests

- coverage for tests can be output by running tests with the `-cover` flag:

    ```bash
    $ go test -cover
    ```

