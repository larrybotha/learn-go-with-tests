# Sync

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Resources](#resources)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- `sync.WaitGroup` allows one to run processes concurrently, and to block
    execution of subsequent commands until all goroutines are complete
    - an instance of `WaitGroup` is initialised without explicit variable
        assignment:

        ```golang
        var wg sync.WaitGroup
        // expect 5 processes to be run concurrently with this WaitGroup
        wg.Add(5)

        // do something that runs concurrently
        ```
- `sync.WaitGroup.Add()` indicates to a `WaitGroup` how many times a process will
    be run before it is done
- `sync.WaitGroup` works similarly to async / await in Javascript, while
    allowing one to specify the number of processes that will be run, similar to
    some Javascript testing libraries that require one to define upfront how
    many assertions there are in a single test
- `sync.WaitGroup.Wait()` is akin to `await somePromise` in Javascript - it
    blocks further processing until all processes are complete
- `sync.WaitGroup.Done()` is used to indicate that a specific process is
    complete. This is similar to `Promise.resolve` in Javascript. e.g.

    ```golang
    var wg sync.WaitGroup
    wg.Add(5)

    for i := 0; i < 5; i++ {
      go func(wg *sync.WaitGroup) {
        // do something
        wg.Done()
      }(&wg)
    }

    wg.Wait()
    ```
- a mutex is a `Mutual Exclusion Lock` - it allows one to lock mutations of
    instances when the instance may be being updated concurrently, and then
    unlock the instance when execution is complete, allowing other processes to
    continue processing
- to use a mutex:
    - define a mutex on the instance's struct using `sync.Mutex`:

        ```golang
        type MyStruct struc {
          mu sync.Mutex
        }
        ```
    - inside methods that mutate properties of the struct, lock the instance
        using the mutex, complete processing, and then unlock the instance:

        ```golang
        func (m *MyStruct) MyMethod() {
          m.mu.Lock()
          // unlock once processing is complete
          defer m.mu.Unlock()

          // perform mutations to instance
        }
        ```
- using `go vet` can reveal issues with one's packages that don't prevent
    compilation. e.g., structs that contain mutexes should only ever be passed
    around by reference / as pointers, as a mutex should not be copied once
    instantiated.
- if one is creating a struct that should only be used as a pointer, e.g. in the
    case where it contains a mutex, one can provide a convenience function to
    create the struct as a pointer, instead of leaving it up to the user:

    ```golang
    type MyStruct struct {
      mu sync.Mutex
    }

    func NewMyStruct() *MyStruct {
      return &MyStruct{}
    }
    ```
- mutexes and channels both allow for concurrent processing. As a general rule:
    - use a mutex when managing internal state of instances of values
    - use channels when passing ownership of data

### Resources

- [Mutex or Channel](https://github.com/golang/go/wiki/MutexOrChannel)
