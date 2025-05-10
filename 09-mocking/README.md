# 9 Mocking

- mocking in Go is useful as one can test behaviour of interfaces,
    as long as the struct used for testing implements the interface
- iterators in Go are higher-order functions that return a `iter.Seq[Type]`
    function. This function accepts a `yield` argument which is a function
    that accepts some value, and returns a `bool`
    * iteration can be stopped at any point by returning `false`
    * the `yield` function in the iterator can be envisioned as a single
        iteration of a loop. The body of the loop would represent the
        body of the `yield`, while the value passed to the `yield` would
        be the specific iterated value in the `range` statement.

        For ranges where there is a key-value pair, `iter.Seq2` would be
        used

        A `false` return from the `yield` indicates the end of the
        sequence, but what does that mean in terms of the body of a
        loop?

        Either an explicit `break`, or the loop completing constitutes
        a `false` return value. `continue` is equivalent to a `true`
        result in the call to `yield`

## links and resources

- https://medium.com/eureka-engineering/a-look-at-iterators-in-go-f8e86062937c





