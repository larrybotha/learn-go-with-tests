# 8 Dependency injection

- using interfaces as function parameters allows for one to decouple a function
    from how the value is used within the function
    * e.g. we want to _write_ a string somewhere after processing it, but we don't
        want to make an assumption as to where it's written

        i.e. to a HTTP response, a file, a buffer, etc.

        To do this, a function can expect a parameter that implements `io.Writer`
        which would be a `struct` that determines _where_ the string would be
        written to

