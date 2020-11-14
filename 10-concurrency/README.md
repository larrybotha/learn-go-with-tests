# Concurrency

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

- Go achieves concurrency through the use of goroutines and channels
- a goroutine is a function that is called when it is preceded with `go`
    - a goroutine is a non-blocking function call
    - as in Javascript, anonymous functions can be immediately invoked:

        ```golang
        func main() {
          go func() {
            time.Sleep(time.Second)

            fmt.Println("sleep over")
          }()

          fmt.Println("will print before sleep")
          time.Sleep(2 * time.Second)
        }
        ```
- a channel can be created using `make`
    ```golang
    type chanType struct {
      string
      int
    }

    myChannel = make(chan chanType)
    ```
- using a channel requires 2 separate steps:
    - a _send statement_
      - this sends a value into a channel:

          ```golang
          type chanType struct{
            string
            int
          }

          func main() {
            myChan := make(chan chanType)

            // a send statement
            myChan <- chanType{"foo", 1}
          }
          ```
    - a _receive expression_, where a value is extracted from a channel:

        ```golang
        type chanType struct {
          string
          int
        }

        func main() {
          myChan := make(chan chanType)

          go func() {
            // send statement
            myChan <- chanType{"foo", 1}
          }()

          // receive expression
          result := <-myChan

          fmt.Printf("%v", result)
        }
        ```
- is with closures and `this` in Javascript, iteration and closures with send
    statements on channels can result in odd results unless one passes the
    currently iterated value into the closure when it is called:

    ```golang
    type chanType struct {
      int
    }

    func main() {
      myIncorrectChan := make(chan chanType)
      myCorrectChan := make(chan chanType)
      xs := []int{1, 2, 3, 4}

      for _, num := range xs {
        // sends `4` to channel every iteration
        go func() {
          myIncorrectChan <- chanType{num}
        }()

        // sends subsequent slice value to channel every iteration
        go func(n int) {
          myCorrectChan <- chanType{num}
        }(num)
      }

      for i := 0; i < len(xs); i++ {
        result := <-myIncorrectChan
        fmt.Printf("incorrect value %d\n", result)
      }

      for i := 0; i < len(xs); i++ {
        result := <-myCorrectChan
        fmt.Printf("correct value %d\n", result)
      }
    }
    ```
- one can have Go indicate race errors by running tests with `-race`

### Tests

- to simulate a slow request or execution of a function, one can use a stub with
    `time.Sleep`:

    ```golang
    func mySlowFunctionStub(_ someType) {
      time.Sleep(1 * time.Millisecond)
    }
    ```

## Resources

