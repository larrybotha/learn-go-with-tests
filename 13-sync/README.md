# 13 Sync

- mutexes are an additional tool in Go for concurrency
- mutexes are suited to concurrency where state is managed
- channels are suited to transferring ownership of data
- `sync.WaitGroup` is useful when a known number of goroutines is going to be created,
    we need to wait synchronously for all of them to be done
    * `WaitGroup.Add` adds to the number of items that should be waited for in a
        `WaitGroup`
    * `WaitGroup.Done` should be called inside a goroutine once the task is complete.
        It decrements the counter of the `WaitGroup`
    * `WaitGroup.Wait` will block code until the counter hits zero
- structs that contain `sync.Mutex` should be passed around as a pointer. Copying a `Mutex`
    should be avoided
    * a copy of a `Mutex` means we have a _new_ concurrent context - locking the copied
        mutex will mean the values we're trying to prevent from being modified will
        be modifiable
    * to encourage the use of the struct as a pointer, one can add a `New[MyStruct]`
        function to a package, returning the value in the safe form by default
- `go vet` can pick up code where a `Mutex` will be copied - passing a pointer to the
    value resolves the issue

## links and resources

- https://go.dev/wiki/MutexOrChannel




