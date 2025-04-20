# Sync

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- context in Go allows one to share data between calls, much in the same way as
    in React and Svelte
- to create a new context, one may use:
    - `context.Context()` - creates a context
    - `context.WithCancel(ctx)` - create a copy of `ctx` that allows for cancelling
        contextj
    - `context.WithDeadline(ctx, time)` - create a copy of `ctx` that will be done if
        the time passed in is exceeded
    - `context.WithTimeout(ctx, duration)` - create a copy of `ctx` that will be
        done after a specified duration
- a context's `Done()` method returns a channel that will receive a value when
    the context errors, its deadline is passed, it's timeout is exceeded, or it
    is cancelled
- it is discouraged to pass values around using context and `context.Value`, as
    the value contained in the context is not typed

### Tests

- one can test that a request was cancelled by ensuring that data was not
    written to the response
    - unfortunately, `httptest.ResponseRecorder`, which is returned after calling
        `httptest.NewRecorder()`, does not allow for one to evaluate whether data
        was written to the response or not. `httptest.ResponseRecorder` is an
        implementation of `http.ResponseWriter` that records mutations for
        later inspection during tests
    - to address this, we can create a spy that implements
        `httptest.ResponseRecorder`

        ```golang
        type ResponseWriterSpy struct {
            written bool
        }

        func (r *ResponseWriterSpy) Header() http.Header {
            r.written = true
            return nil
        }

        func (r *ResponseWriterSpy) Write([]byte) (int, err) {
            r.written = true
            return 0, errors.New("not implemented")
        }

        func (r *ResponseWriterSpy) WriteHeader(statusCode int) {
            r.written = true
        }
        ```
    - with this spy, we can evaluate `ResponseWriterSpy.written` to determine
        whether a request ever resulted in data being written
