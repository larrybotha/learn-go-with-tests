# 11 Select

- `select` is like a `switch` statement for channels
- the first channel to return a result will short-circuit the statement
- when using a `select` statement, it's important to consider scenarios where
    it may never terminate. In these cases, a timeout can be useful, using
    `time.After`, which is a channel that will return after a given duration

