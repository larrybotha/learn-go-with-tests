# Select

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- `select` behaves in a similar manner for channels as Promise.race does for
    promises in Javascript
- `select` looks like a switch statement, except that each case receives a value
    from a channel:

    ```golang
    select {
      case <-chanA:
        // do something
      case <-chanB:
        // do something
    }
    ```
- `timer.After` creates a channel that will receive a value after a provided
    duration
- `time.After` can be used in a `select` statement to handle timeouts, or
    channels never receiving a value:

    ```golang
    select {
      case <-longRunningChannel
        // do something
      case <- time.After(10 * time.Millisecond)
        // do something
    }
    ```
- a channel can be manually closed using the `close` builtin function:

    ```golang
    func createUrlResolvedChannel(url string) {
      c := make(c struct{})

      go func() {
        http.Get(url)
        close(c)
      }()

      return c
    }
    ```

### Tests

- use `t.Fatal` to prevent a test from any further evaluation if there are other
    possible assertions that could be evaluated after the current one and it
    fails
- `http.httptest` can be used to create servers and generate URLs for tests

