# 02 - Integers

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

- [spew](https://github.com/davecgh/go-spew) is useful for dumping values
- [testify](https://github.com/stretchr/testify) may be useful for writing
    assertions

### Go

- Go has no `while` or `do` loops - only `for`
- the initial and update statements in a `for` loop are optional
- a `for` loop with only a condition is Go's while loop
- an infinite loop can be created by providing no parameters

### Tests

- benchmarks can be created using `BenchmarkXxx`
- benchmarks accept a parameter of type `testing.B`
- benchmarks should contain a loop inside of which the function under test is
    executed
    - the condition for the loop should use the `b.N` value - a value the
        testing framework uses to determine an acceptable value for how many
        times the benchmark should be run
