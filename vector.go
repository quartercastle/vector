package vector

import (
	"errors"
	"math"
)

// functions from the gonum package that optimizes arithmetic
// operations on lists of float64 values
func axpyUnitaryTo(dst []float64, alpha float64, x, y []float64)
func scalUnitaryTo(dst []float64, alpha float64, x []float64)

// Vector is the definition of a row vector that contains scalars as
// 64 bit floats
type Vector []float64

// axis is an integer enum type that describes vector axis
type axis int

const (
	// the consts below are used to represent vector axis, they are useful
	// to lookup values within the vector.
	x axis = iota
	y
	z
)

var (
	// ErrNot3Dimensional is an error that is returned in functions that only
	// supports 3 dimensional vectors
	ErrNot3Dimensional = errors.New("vector is not 3 dimensional")
)

// Clone a vector
func Clone(v Vector) Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

// Clone a vector
func (v Vector) Clone() Vector {
	return Clone(v)
}

// Add a vector with a vector or a set of vectors
func Add(v1 Vector, vs ...Vector) Vector {
	return v1.Clone().Add(vs...)
}

// Add a vector with a vector or a set of vectors
func (v Vector) Add(vs ...Vector) Vector {
	dimensions := len(v)

	for i := range vs {
		if vd := len(vs[i]); vd > dimensions {
			v = append(v, make(Vector, vd-dimensions)...)
			dimensions += vd - dimensions
		}
		axpyUnitaryTo(v, 1, v, vs[i])
	}

	return v
}

// Sub subtracts a vector with another vector or a set of vectors
func Sub(v1 Vector, vs ...Vector) Vector {
	return v1.Clone().Sub(vs...)
}

// Sub subtracts a vector with another vector or a set of vectors
func (v Vector) Sub(vs ...Vector) Vector {
	dimensions := len(v)

	for i := range vs {
		if vd := len(vs[i]); vd > dimensions {
			v = append(v, make(Vector, vd-dimensions)...)
			dimensions += vd - dimensions
		}
		axpyUnitaryTo(v, -1, vs[i], v)
	}

	return v
}

// Scale vector with a given size
func Scale(v Vector, size float64) Vector {
	return v.Clone().Scale(size)
}

// Scale vector with a given size
func (v Vector) Scale(size float64) Vector {
	scalUnitaryTo(v, size, v)
	return v
}

// Equal compares that two vectors are equal to each other
func Equal(v1, v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if math.Abs(v1[i]-v2[i]) > 1e-5 {
			return false
		}
	}

	return true
}

// Equal compares that two vectors are equal to each other
func (v Vector) Equal(v2 Vector) bool {
	return Equal(v, v2)
}

// Magnitude of a vector
func Magnitude(v Vector) float64 {
	var result float64

	for _, scalar := range v {
		result += math.Pow(scalar, 2)
	}

	return math.Sqrt(result)
}

// Magnitude of a vector
func (v Vector) Magnitude() float64 {
	return Magnitude(v)
}

// Unit returns a direction vector with the length of one.
func Unit(v Vector) Vector {
	return v.Clone().Unit()
}

// Unit returns a direction vector with the length of one.
func (v Vector) Unit() Vector {
	l := v.Magnitude()

	if math.Abs(l) < 1e-8 {
		return v
	}

	for i := range v {
		v[i] = v[i] / l
	}

	return v
}

// Dot product of two vectors
func Dot(v1, v2 Vector) float64 {
	var result float64

	if len(v1) > len(v2) {
		v2 = append(v2, make(Vector, len(v1)-len(v2))...)
	}

	if len(v1) < len(v2) {
		v1 = append(v1, make(Vector, len(v2)-len(v1))...)
	}

	for i := range v1 {
		result += v1[i] * v2[i]
	}

	return result
}

// Dot product of two vectors
func (v Vector) Dot(v2 Vector) float64 {
	return Dot(v, v2)
}

// Cross product of two vectors
func Cross(v1, v2 Vector) (Vector, error) {
	if len(v1) != 3 || len(v2) != 3 {
		return Vector{}, ErrNot3Dimensional
	}

	return Vector{
		v1[y]*v2[z] - v1[z]*v2[y],
		v1[z]*v2[x] - v1[x]*v2[z],
		v1[x]*v2[z] - v1[z]*v2[x],
	}, nil
}

// Cross product of two vectors
func (v Vector) Cross(v2 Vector) (Vector, error) {
	return Cross(v, v2)
}
