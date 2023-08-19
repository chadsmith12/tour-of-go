# Types: Pointers, Structs, Slices, and Maps

## Pointers

go has pointers. A pointer holds lthe memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```
var p *int
// this will make 'p' an int pointer.
```

The `&` operator generates a pointer to its operand.

```
i := 42
// p is a pointer to i
p = %i
```

The `*` operator denotes the pointer's underlying value.

```
fmt.Println(*p) // read i through the pointer p
*p = 21 // set i through the pointer p
```

This is known as "dereferencing" and "indirecting".

## Structs

A `struct` is ac ollection of fields.

### Struct Fields

Struct fields are accessed using a dot.

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex {1 2}
	v.X = 4
}
```

## Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression 

```go
var a [10]int
```

declares a variable `a` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

### Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type '[]T' is a slice with elements of type `T`

A slice is formed by specifying two indices, a low an a high bound, separated by a colon:

```go
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 a `a`.

```go
a[1":4]
```

### Slices are like references to arrays

A slice does not store any data. It just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same u nderlying array will see those changes.

### Slice Literals

A slice literal is like an array literal without the length.

This is an array literal:

```go
[3]bool {true, true, false}
```

And this creates the same array as above, then builds a slice that references it.

```go
[]bool {true, true, false}
```

### Slice length and capacity

A slice has both a lengh and a cpacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend the slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

### Nil slices

The zero value of a slice is `nil`.

A nil slice has a length and capacity of 0 and has no underlying array.

### Appending to a slice

You can use the `append` function to append to a slice. It takes in a slice, and the values you want to append. It returns back the slice. The function will initailize new memory for the array if it has to and return back the new array to you.

---

## Range

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

---

## Maps

A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added. The `make` function returns a map of the given type, initialized and ready for use.

Map literals are like struct literals, but the keys are required.

```go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "Bell Labs": Vertex {
         40.68, -74.39967
    },
    "Google": Vertex {
         37.422, -122.08408
    }
}

```

If the top-level type is just a type name, you can omit it from the elements of the literal.


```go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "Bell Labs": {
         40.68, -74.39967
    },
    "Google": {
         37.422, -122.08408
    }
}

```

### Mutating Maps

Insert or update an element in a map:

```go
m[key] = elem
```

Retrieve an element:

```go
elem = m[key]
```

Delete an element:

```go
delete(m, key)
```

Test that a key is present with a two-value assignment:

```go
elem, ok := m[key]

if ok == true {
     // key inside the map
}
```

---

## Function Values

Functions are values too. They can be passed around just like other values. Function values may be used as function arguments and return values.


```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}
```
### Function Closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the references variables; in this sense the function is "bound" to the variables.

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
	return sum
    }
}
```

