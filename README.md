# vector

[![Version](https://img.shields.io/github/release/kvartborg/vector.svg)](https://github.com/kvartborg/vector/releases)
[![GoDoc](https://godoc.org/github.com/kvartborg/vector?status.svg)](https://pkg.go.dev/github.com/kvartborg/vector?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/kvartborg/vector)](https://goreportcard.com/report/github.com/kvartborg/vector)

The motivation behind this package is to find a better way to write vector math
in Golang, there has to be a more expressive way without it getting to verbose.

This is an experiment, there are lots of other libraries tackling vector math as
well. You should properly take a look at [`gonum`](https://github.com/gonum/gonum) before consider using this package.

## Install
```sh
go get github.com/kvartborg/vector
```

## Usage
Golang does not have a way to define generic types yet, which
limits this package to operate with vectors represented as `float64` values only.
To allow for multi-dimensional vectors, a vector is simply represented as
a list of `float64` values.
```go
// A Vector is simply a list of float64 values
type Vector []float64
```
> I will consider adding `float32` and `int` support at a later stage,
  if there is a good reason or when go adds generics to the language.

### Tackling verbosity
Another goal of this experiment is to minimize the verbosity around using the package,
this can be achieved by using type aliasing. In this way you can omit the package
identifier and give the `Vector` a shorter name like `vec` or something else,
it is up to you.
```go
// Minimize the verbosity by using type aliasing
type vec = vector.Vector

// addition of two vectors
result := vec{1, 2}.Add(vec{2, 4})
```

A nice side effect of representing a vector as a list of `float64` values is that
a slice of `float64` values can easily be turned into a vector by using type casting.
This elimitates the need for any constructor functions for the vector type.
```go
// Turn a list of floats into a vector
v := vec([]float64{1, 2, 3})
```

### Mutability vs Immutability
All arithmetic operations are immutable by default. But if needed a `Vector` can be
turned into a `MutableVector` with the `vector.In` function, see example below.
A mutable vector performs arithemetic operations much faster without taking up
any memory.
```go
// create vectors
v1, v2 := vec{1, 2}, vec{2, 4}

// Immutable addition, will return a new vector containing the result.
result := v1.Add(v2)

// Mutable addition, will do the calculation in place in the v1 vector
vector.In(v1).Add(v2)
```

### Slicing a vector
Another benefit of using a list of `float64` to represent a vector is that you
can slice vectors as you normally would slice lists in go.
```go
v1 := vec{1, 2, 3}
v2 := v1[1:] // returns a new vec{2, 3}
```

## Documentation
The full documentation of the package can be found on [godoc](https://pkg.go.dev/github.com/kvartborg/vector?tab=doc).

## Contributions
Contributions with common vector operations that are not included in this package are welcome.

## Credits
Thanks to [`gonum`](https://github.com/gonum/gonum) for inspiration and the following functions [`axpyUnitaryTo`](https://github.com/gonum/gonum/blob/master/internal/asm/f64/axpyunitaryto_amd64.s), [`scalUnitaryTo`](https://github.com/gonum/gonum/blob/c3867503e73e5c3fee7ab93e3c2c562eb2be8178/internal/asm/f64/scalunitaryto_amd64.s) that enhances the performance of arithmetic operations in this package.

## License
This project is licensed under the [MIT License](https://github.com/kvartborg/vector/blob/master/LICENSE) and includes [`gonum`](https://github.com/gonum/gonum) code that is licensed under [3-Clause BSD license](https://github.com/gonum/gonum/blob/master/LICENSE).
