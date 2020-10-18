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

- structs define the types of entities, and follow the pattern:

    ```golang
    type SomeObject struct {
      Property1 [type]
      Property2 [type]
    }
    ```
- values defined by structs can be instantiated in one of 2 ways:
  - implicit - property names are not provided, and are instantiated in the
      order defined in the struct:

      ```golang
      type Animal struct {
        Legs int
        Eyes int
      }

      cat := Animal{4, 2}
      ```
  - explicit - property names are provided during instantiation:

      ```golang
      type Animal struct {
        Legs int
        Eyes int
      }

      cat := Animal{Legs: 4, Eyes: 2}
      ```
- properties of structs are accessed by name:

    ```golang
    type Animal struct {
      Legs int
      Eyes int
    }

    cat := Animal{Legs: 4, Eyes: 2}

    fmt.Printf("%d", cat.Legs)
    ```
- methods on structs are defined at the location of the function definition, and
    follow this pattern:

    ```golang
    func (s StructName) MethodName(args ...) {}
    ```

    e.g. for `Animal`:

    ```golang
    func (a Animal) GetAge() int {
      return a.Age;
    }
    ```
- the convention for defining methods is to use the lowercase first letter of
    the struct as the name for the argument in the context section of the
    function definition (i.e. where the name of the struct is used)
- the argument to the context section of the method definition is a reference to
    the instance of the struct, which can be used inside the function body. This
    is equivalent to `this` in other programming languages
- for structs that share similar methods, an interface can be used to define how
    the method or properties should be defined
    - unlike languages where it's common to use a `MyType implements
        MyInterface` syntax, in Go, interface implementation is implicit

        ```golang
        // before - Area is a method on each struct, but with no relation
        type Apple struct {
          Vitamins int
        }

        func (a Apple) Nutrition() int {}

        type Beef struct {
          Proteins int
        }

        func (b Beef) Nutrition() int {}


        // after - we use an interface to create a common definition for food.
        // Now, if calculating the nutrition of a food, we can reference the
        // type of a variable as Food
        type Food interface {
          Nutrition() int
        }

        type Apple struct {
          Vitamins int
        }

        func (a Apple) Nutrition() int {}

        type Beef struct {
          Proteins int
        }

        func (b Beef) Nutrition() int {}
        ```
- the syntax for `range` in `for` loops is:

    ```golang
    for index, value := range xs {
      //
    }
    ```
- one can define anonymous structs during variable instantiation:

    ```golang
    myVar := struct {
      name string
      age  int
    } {
      name: "Joe",
      age:  5
    }

    // or as a slice
    xs := []struct {
      foo string
      bar int
    } {
      {foo: "a", bar: 1},
      {foo: "b", bar: 2},
    }
    ```
- `%v` will print the value of a variable
  - `%+v` will print the value with field names
  - `%#v` will print the value using Go's formatting
- `%g` will print decimal numbers

### Tests

- Go allows for parameterised tests to be run using a simple `for` loop:

    ```golang
    func TestMyMethod(t *testing.T) {
      tests := []struct {
        name  string
        value MyInterface
        want  int
      } {
        {name: "test 1", value: StructA{1}, want: 1},
        {name: "test 2", value: StructB{2}, want: 2},
        {name: "test 3", value: StructA{3}, want: 3},
      }

      for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
          got := tt.value.MyMethod()

          if got != tt.want {
            t.Errorf("%#v: got %d, wanted %d", tt.value, got, tt.want)
          }
        })
      }
    }
    ```
