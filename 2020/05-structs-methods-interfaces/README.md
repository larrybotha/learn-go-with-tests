# Structs, methods, and interfaces

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Takeaways](#takeaways)
  - [Go](#go)
  - [Tests](#tests)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Takeaways

### Go

- structs define the shape of an object:

    ```golang
    type MyStruct struct {
      privateProp     [type]
      PublicProp      [type]
      privateMethod() [type]
      PublicMethod()  [type]
    }
    ```
- pulic methods and properties on structs are defined with an uppercase first
    letter
- private methods and properties on structs are defined with a lowercase first
    letter
- interfaces are not explicitly extended as in other languages - one simply
    needs to implement the method in an interface
    - this allows one to implement methods from existing libraries, such as
        implementing `String()` on the `Stringer` interface for custom string
        format output
-  methods on structs have a receiver function that where the name of the
    instance is usually just the first letter of the type. This value is similar
    to `this` in other languages
- methods are defined in the following way:

    ```golang
    func (m MyType) MyMethod() {}
    ```
- the `%.{d}f` formatting verb will print a floating point value to `d` decimal values
- the `%g` formatting will print the full floating point value
- anonymous structs are structs defined inline without names:

    ```golang
    myValue := struct{
      name  string
      age   int
    } {"Joe", 21}

    mySlice := []struct {
      name  string
      age   int
    } {
      {"Jane", 22},
      {"Charlie", 32},
    }
    ```

### Tests

- Go allows for parameterised tests to be run by simply iterating over a slice:

    ```golang
    func TestSomething(t *testing.T) {
      tests := []struct {
        name  string
        value [type]
        want  [type]
      }{
        {name: "test 1 name", value: "foo", want: "bar"},
        {name: "test 2 name", value: "baz", want: "quux"},
      }

      for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
          got := Something(tt.value)

          if got != tt.want {
            t.Errorf("got %d, want %d", got, tt.want)
          }
        })
      }
    }
    ```
