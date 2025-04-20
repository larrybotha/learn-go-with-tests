# Dependency Injection

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

- functions that return multiple values define those values in parenthesis:

    ```golang
    func MyFunc() (n int, s string, xs []int) {}
    ```
- to implement an interface on a struct, the method must be implemented with a
    receiver:

    ```golang
    type MyStruct struct {}

    // MyStruct implements Write on the io.Writer interface
    func (m MyStruct) Write(p byte[]) (n int, err error) {}
    ```
- like `sleep` in bash, Go has a `time.Sleep` method which accepts a value of
    type `time.Duration`:

    ```golang
    var sleepTime time.Duration = 5 * time.Second
    time.Sleep(sleepTime)
    ```
- one can use a struct to encapsulate configurations:

    ```golang
    type CustomSleeper struct {
      duration time.Duration
      sleep func(time.Duration)
    }

    func (c CustomSleeper) sleep() {
      c.sleep(c.duration)
    }

    func main() {
      sleeper := CustomSleeper{duration: 5, sleep: time.Sleep}
    }
    ```
- to implement an interface method, one requires a struct on which to implement the
    method
- an interface method can't be defined on an anonymous struct, as interface
    methods have to be defined using receivers, while functions on structs are
    equivalent to assigning a function to a variable
- one can either explicitly pass a value by its address to a function, or
    initialise the variable with the address:

    ```golang
    buffer := bytes.Buffer{}
    fmt.Fprint(&buffer, "foo")

    // or
    bufferAddress = &bytes.Buffer{}
    fmt.Fprint(bufferAddress, "foo")
    ```
- a single struct may implement multiple interfaces

### Tests

- a simple mock can be implemented with a struct that mutates an internal
    property each time the interface it implemnts is called:

    ```golang
    type MySpy struct {
      Calls []string
    }

    // create a mock for time.Sleep
    func (m *MySpy) Sleep(duration time.Duration) {
      m.Calls = append(m.Calls, "called")
    }
    ```


## Resources

- [Uncle Bob's "When to mock"](https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html)
- [The little mocker](https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html)
    - A conversation covering the differences between mocks, stubs, spies, and
        other test doubles
