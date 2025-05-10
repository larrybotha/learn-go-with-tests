# 10 Concurrency

- channels can be sent values, and channels can emit values
- to synchronously wait for all values in a channel to respond:
    * loop over a range of the expected number of items in the channel
    * aggregate the values received from the channel
    * this is _somewhat_ like using `sync.WaitGroup`
- `testing.B.ResetTimer` is used in benchmarks to reset the elapsed time,
    among other things
- `testing.B.Loop` is syntactic sugar for the legacy `b.N` that was used
    to determine how many iterations there would be in the benchmark

