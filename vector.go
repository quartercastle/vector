package vector

import (
	"errors"
	"math"
)

// Vector is the definition of a row vector that contains scalars as
// 64 bit floats
type Vector []float64

type axis int

const (
	x axis = iota
	y
	z
)

var (
	ErrNot3Dimensional = errors.New("vector is not 3 dimensional")
)

func Clone(v Vector) Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

func (v Vector) Clone() Vector {
	return Clone(v)
}

func Add(v1 Vector, vs ...Vector) Vector {
	return v1.Clone().Add(vs...)
}

func (v1 Vector) Add(vs ...Vector) Vector {
	dimensions := len(v1)

	for _, v := range vs {
		for i := range v {
			if i >= dimensions {
				ed := len(v) - dimensions
				v1 = append(v1, make(Vector, ed)...)
				dimensions += ed
			}
			v1[i] += v[i]
		}
	}

	return v1
}

func Sub(v1 Vector, vs ...Vector) Vector {
	return v1.Clone().Sub(vs...)
}

func (v1 Vector) Sub(vs ...Vector) Vector {
	dimensions := len(v1)

	for _, v := range vs {
		for i := range v {
			if i >= dimensions {
				ed := len(v) - dimensions
				v1 = append(v1, make(Vector, ed)...)
				dimensions += ed
			}
			v1[i] -= v[i]
		}
	}

	return v1
}

func Scale(v Vector, size float64) Vector {
	result := v.Clone()

	for i := range v {
		result[i] *= size
	}

	return result
}

func (v Vector) Scale(size float64) Vector {
	return Scale(v, size)
}

func Equals(v1, v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if !(math.Abs(v1[i]-v2[i]) < 1e-5) {
			return false
		}
	}

	return true
}

func (v1 Vector) Equals(v2 Vector) bool {
	return Equals(v1, v2)
}

func Magnitude(v Vector) float64 {
	var result float64

	for _, scalar := range v {
		result += math.Pow(scalar, 2)
	}

	return math.Sqrt(result)
}

func (v Vector) Magnitude() float64 {
	return Magnitude(v)
}

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

func (v1 Vector) Dot(v2 Vector) float64 {
	return Dot(v1, v2)
}

func Cross(v1, v2 Vector) Vector {
	if len(v1) != 3 || len(v2) != 3 {
		panic(ErrNot3Dimensional)
	}

	return Vector{
		v1[y]*v2[z] - v1[z]*v2[y],
		v1[z]*v2[x] - v1[x]*v2[z],
		v1[x]*v2[z] - v1[z]*v2[x],
	}
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Cross(v1, v2)
}
