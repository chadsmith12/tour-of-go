# Methods and Interfaces

## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument. The receiver appears in its own argument list between `func` keyword and the method name. In the example below, the `Abs` method has a receiver of type `Vertex` named `v`.

```go
type Vertex struct {
     X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
```

You can declare a method on non-struct types, too. You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes built-in types such as `int`).

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }

    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
```

### Pointer Receivers

You can declare methods with pointer receivers. This means the receiver type has the literal syntax `*T` for some type `T`.

For example: We could create a `Scale` method defined on the `*Vertex`. Methods with pointer reveivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
If you make a `Scale` method with just a value receiver then notice that nothing mutates.

You can write the methods as just regular functions that take in the type instead of putting them directly on the type.

### Choosing value or pointer receiver

There are two reasons to use a pointer receiver:

1: So that the method can modify the value that its receiver points to
2: Avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

### Interfaces

An **interface type** is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.

Types will implement interfaces by implementing its methods. There is not explicit declaration of intent and not implements keyword. 

Implicit interfaces decouple the definition of an interface from its implementation, which would then appaer in any package without prearrangement.

### Interface Values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type.

### Nil interface values

A nil interface holds neither value nor concrete type. Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

### The empty interface

The interface type that specifies zero methods is known as the **empty interface**:

```go
interface{}
```

An empty interface may hold values of any type. Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

### Type assertions

A ** type assertion ** provides access to an interface value's unerlying concrete value. 

```go
t := i(T)
```
This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`. If it does not hold a `T`, then the statement will trigger a panic.

The test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```go
t, ok := i.(T)
```

If `i` holds a `T` then `t` will be the underlying value and `ok` will be true. If not then `ok` will be false and `t` will be the zero value of type `T`, and not panic occurs.

### Type switches

A type switch is a construct that permits several type assertions in series. A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```go
switch v := i.(type) {
    case T:
        // here v as type T
    case S:
        // here v has type S
    default:
        // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

### Stringers

One of the most ubiquitous interfaces is `Stringer` defiend by the `fmt` package. A `Stringer` is a type that can describe itself as a string. Packages look for this interface to print values.

```go
type Stringer interface {
    String() string
}
```

### Errors

Go programs express error state with `error` values. The `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

Functions often returns an `error` value, and calling code should handle errors by testing whether the error equals `nil`. A nil `error` denotes success; a non-nil error denotres Failure.

## Readers

The `io` package specifies the `io.Reader` interface, which represents the read end of a stread of data. The Go standard library contains many implementations of this interface, including files, network connections, compresssors, ciphers, and others. 

The `io.Reader` interface has a Read method:
```go
func (T) Read(b []bytes) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

