# vector

[![Version](https://img.shields.io/github/release/kvartborg/vector.svg)](https://github.com/kvartborg/vector/releases)
[![Build Status](https://travis-ci.org/kvartborg/vector.svg?branch=master)](https://travis-ci.org/kvartborg/vector)
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
we easily can turn a slice of `float64` into a vector by using type casting.
This elimitates the need for any constructor functions for the vector type.
```go
// Turn a list of floats into a vector
v := vec([]float64{1, 2, 3})
```

### Mutability vs Immutability
Most of the arithmetic operations provided by this package has a mutable and immutable implementation.
The way this is separated is that all package level functions are immutable
and methods called on the vector it self are mutable on the calling vector.
```go
// create vectors
v1, v2 := vec{1, 2}, vec{2, 4}

// Immutable addition, returns a new vector
result := vector.Add(v1, v2)

// Mutable addition, will do the calculation in place in the v1 vector
result := v1.Add(v2)
```

The mutable implementation is a lot faster and uses less memory compared to the immutable, this is because the calculation is done in place.
Mutable operations shall be used with care, because it can lead to bugs that can be hard to spot. A use case i find useful is to use the mutable operations when you are inlining the instantiation of the vectors.

```go
// example of safe usage of mutable operations when inlining vector instantiation.
result := vec{1, 2}.Add(vec{2, 4})
```

Or if you are creating a receiving vector where the calculation can be done in place.

```go
// Create result vector where the calculation can be done in place.
result := make(vec, 2)

// Create vectors
v1, v2 := vec{1, 2}, vec{2, 4}

// safe usage of mutable operation when called on a result vector
result.Add(v1, v2)
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
