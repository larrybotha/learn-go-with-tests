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
- an array defined with `[...]int{}` tells the compiler to count the length for
    you
- a slice in Go has no defined length, and can be defined by omitting the array
    parameter when defining the value:

    ```golang
    xs = []int{1,2,3,4,5}
    ```
- a slice can also be created using the built-in `make` function:

    ```golang
    xs := make([]int, len, cap)
    ```

    - `make`s first argument is the type of slice to create
    - `len` is the number of items to initialise the slice with
    - `cap` is the capacity of the slice - if an item is assigned to an index
        beyond the capacity of a slice, a runtime error will be thrown
- a slice can be extended using the built-in `append` function:

    ```golang
    xs := make([]int, 1, 1)
    xs[0] = 1

    xs = append(xs, 2, 3)
    ```

    - `append` returns a copy of the slice passed in, not a reference to the
        original
    - `append` is similar to `Array.prototype.concat` in Javascript
- one can slice a slice using the `xs[low:high]` syntax
    - the slice syntax creates a reference to the original slice - changes to
        the new slice will effect the original slice
    - this is similar to Javascript's `Array.prototype.slice`, but with a
        reference instead of a new value
    - slicing a slice beyond its capacity will throw a runtime error - beware of
        this!
- one can spread values using th `...` syntax:

    ```golang
    xs := []int{1}
    ys := []int{2,3}

    xs = append(xs, ys...)
    ```
- the rest parameter in Go precedes the type, after the value

    ```golang
    func myFunc(xs ...int) {
      // whole bunch numbers
    }
    ```

    - this is similar to Javascript's rest, but remember - it comes AFTER the
        value
- `for` loops allow for the use of a `range` statement that will iterate over an
    array or slice, returning an index and value for each iteration:

    ```golang
    xs = []int{1,2,3,4,5}

    for i, number := range {
      fmt.Printf("%d => %d", i, number)
    }
    ```
- attempting to assign values to indices outside of the capacity of a slice will
    throw runtime errors - DO NOT FORGET THIS
- `reflect.DeepEqual` can be used to check equality of values, but it is not
    typesafe - one can evaluate two values of any two different types, and the
    compiler will not intervene. Use with caution


### Tests

- coverage for tests can be output by running tests with the `-cover` flag:

    ```bash
    $ go test -cover
    ```
- when creating a helper function in one's tests, pass in the instance of
    `*testing.T` of the specific test using that helper
