package vector

import (
	"errors"
	"fmt"
	"math"
)

// Vector is the definition of a row vector that contains scalars as
// 64 bit floats
type Vector []float64

const (
	// the consts below are used to represent vector axis, they are useful
	// to lookup values within the vector.
	x int = iota
	y
	z
)

var (
	// ErrNot3Dimensional is an error that is returned in functions that only
	// supports 3 dimensional vectors
	ErrNot3Dimensional = errors.New("vector is not 3 dimensional")
	// ErrNotSameDimensions is an error that is returned when functions need both
	// Vectors provided to be the same dimensionally
	ErrNotSameDimensions = errors.New("the two vectors provided aren't the same dimensional size")

	// X is the vector axis which can be used to rotate another vector around
	X = Vector{1, 0, 0}
	// Y is the vector axis which can be used to rotate another vector around
	Y = Vector{0, 1, 0}
	// Z is the vector axis which can be used to rotate another vector around
	Z = Vector{0, 0, 1}
)

// Clone a vector
func Clone(v Vector) Vector {
	return v.Clone()
}

// Clone a vector
func (v Vector) Clone() Vector {
	clone := make(Vector, len(v))
	copy(clone, v)
	return clone
}

func Add(v1, v2 Vector) Vector {
	return v1.Clone().Add(v2)
}

func (v Vector) Add(v2 Vector) Vector {
	dim1, dim2 := len(v), len(v2)

	if (dim1 == 1 || dim1 == 2 || dim2 == 3) && dim2 == 1 {
		v[x] += v2[x]
		return v
	}

	if dim1 == 2 && dim2 == 2 {
		v[x], v[y] = v[x]+v2[x], v[y]+v2[y]
		return v
	}

	if dim1 == 3 && dim2 == 2 {
		v[x], v[y] = v[x]+v2[x], v[y]+v2[y]
		return v
	}

	if dim1 == 3 && dim2 == 3 {
		v[x], v[y], v[z] = v[x]+v2[x], v[y]+v2[y], v[z]+v2[z]
		return v
	}

	if dim2 > dim1 {
		axpyUnitaryTo(v, 1, v, v2[:dim1])
	} else {
		axpyUnitaryTo(v, 1, v, v2)
	}

	return v
}

// Sum a vector with a vector or a set of vectors
func Sum(v1 Vector, vs ...Vector) Vector {
	return v1.Clone().Sum(vs...)
}

// Sum a vector with a vector or a set of vectors
func (v Vector) Sum(vs ...Vector) Vector {
	dim := len(v)

	if dim == 2 && len(vs) == 1 && len(vs[0]) == 1 {
		v[x] += vs[0][x]
		return v
	}

	if dim == 2 && len(vs) == 1 && len(vs[0]) == 2 {
		v[x], v[y] = v[x]+vs[0][x], v[y]+vs[0][y]
		return v
	}

	if dim == 3 && len(vs) == 1 && len(vs[0]) == 1 {
		v[x] += vs[0][x]
		return v
	}

	if dim == 3 && len(vs) == 1 && len(vs[0]) == 2 {
		v[x], v[y] = v[x]+vs[0][x], v[y]+vs[0][y]
		return v
	}

	if dim == 3 && len(vs) == 1 && len(vs[0]) == 3 {
		v[x], v[y], v[z] = v[x]+vs[0][x], v[y]+vs[0][y], v[z]+vs[0][z]
		return v
	}

	for i := range vs {
		if len(vs[i]) > dim {
			axpyUnitaryTo(v, 1, v, vs[i][:dim])
		} else {
			axpyUnitaryTo(v, 1, v, vs[i])
		}
	}

	return v
}

// Sub subtracts a vector with another vector or a set of vectors
func Sub(v1 Vector, v2 Vector) Vector {
	return v1.Clone().Sub(v2)
}

// Sub subtracts a vector with another vector or a set of vectors
func (v Vector) Sub(v2 Vector) Vector {
	dim1, dim2 := len(v), len(v2)

	if (dim1 == 1 || dim1 == 2 || dim2 == 3) && len(v2) == 1 {
		v[x] -= v2[x]
		return v
	}

	if dim1 == 2 && dim2 == 2 {
		v[x], v[y] = v[x]-v2[x], v[y]-v2[y]
		return v
	}

	if dim1 == 3 && dim2 == 1 {
		v[x] -= v2[x]
		return v
	}

	if dim1 == 3 && dim2 == 2 {
		v[x], v[y] = v[x]-v2[x], v[y]-v2[y]
		return v
	}

	if dim1 == 3 && dim2 == 3 {
		v[x], v[y], v[z] = v[x]-v2[x], v[y]-v2[y], v[z]-v2[z]
		return v
	}

	if dim2 > dim1 {
		axpyUnitaryTo(v, -1, v2[:dim1], v)
	} else {
		axpyUnitaryTo(v, -1, v2, v)
	}

	return v
}

// Invert returns a copy of the inverse of the provided Vector
func Invert(v Vector) Vector {
	return v.Clone().Invert()
}

// Invert inverts the vector, and then returns it
func (v Vector) Invert() Vector {
	for i := range v {
		v[i] *= -1
	}
	return v
}

// Scale vector with a given size
func Scale(v Vector, size float64) Vector {
	return v.Clone().Scale(size)
}

// Scale vector with a given size
func (v Vector) Scale(size float64) Vector {
	dim := len(v)

	if dim == 2 {
		v[x], v[y] = v[x]*size, v[y]*size
		return v
	}

	if dim == 3 {
		v[x], v[y], v[z] = v[x]*size, v[y]*size, v[z]*size
		return v
	}

	scalUnitaryTo(v, size, v)
	return v
}

// Equal compares that two vectors are equal to each other
func Equal(v1, v2 Vector) bool {
	return v1.Equal(v2)
}

// Equal compares that two vectors are equal to each other
func (v Vector) Equal(v2 Vector) bool {
	dim := len(v)
	if dim != len(v2) {
		return false
	}

	if dim == 2 {
		return math.Abs(v[x]-v2[x]) < 1e-8 && math.Abs(v[y]-v2[y]) < 1e-8
	}

	if dim == 3 {
		return math.Abs(v[x]-v2[x]) < 1e-8 && math.Abs(v[y]-v2[y]) < 1e-8 && math.Abs(v[z]-v2[z]) < 1e-8
	}

	for i := range v {
		if math.Abs(v[i]-v2[i]) > 1e-8 {
			return false
		}
	}

	return true
}

// Magnitude of a vector
func Magnitude(v Vector) float64 {
	return v.Magnitude()
}

// Magnitude of a vector
func (v Vector) Magnitude() float64 {
	dim := len(v)

	if dim == 1 {
		return math.Sqrt(v[x] * v[x])
	}

	if dim == 2 {
		return math.Sqrt(v[x]*v[x] + v[y]*v[y])
	}

	if dim == 3 {
		return math.Sqrt(v[x]*v[x] + v[y]*v[y] + v[z]*v[z])
	}

	var result float64
	for _, scalar := range v {
		result += scalar * scalar
	}

	return math.Sqrt(result)
}

// Unit returns a direction vector with the length of one.
func Unit(v Vector) Vector {
	return v.Clone().Unit()
}

// Unit returns a direction vector with the length of one.
func (v Vector) Unit() Vector {
	dim := len(v)

	if dim == 2 {
		l := math.Sqrt(v[x]*v[x] + v[y]*v[y])
		if l < 1e-8 {
			return v
		}
		v[x], v[y] = v[x]/l, v[y]/l
		return v
	}

	if dim == 3 {
		l := math.Sqrt(v[x]*v[x] + v[y]*v[y] + v[z]*v[z])
		if l < 1e-8 {
			return v
		}
		v[x], v[y], v[z] = v[x]/l, v[y]/l, v[z]/l
		return v
	}

	l := v.Magnitude()

	if l < 1e-8 {
		return v
	}

	for i := range v {
		v[i] = v[i] / l
	}

	return v
}

// Dot product of two vectors
func Dot(v1, v2 Vector) float64 {
	result, dim1, dim2 := 0., len(v1), len(v2)

	if dim1 > dim2 {
		v2 = append(v2, make(Vector, dim1-dim2)...)
	}

	if dim1 < dim2 {
		v1 = append(v1, make(Vector, dim2-dim1)...)
	}

	if dim1 == 2 {
		return v1[x]*v2[x] + v1[y]*v2[y]
	}

	if dim1 == 3 {
		return v1[x]*v2[x] + v1[y]*v2[y] + v1[z]*v2[z]
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
	return v1.Cross(v2)
}

// Cross product of two vectors
func (v Vector) Cross(v2 Vector) (Vector, error) {
	if len(v) != 3 || len(v2) != 3 {
		return nil, ErrNot3Dimensional
	}

	return Vector{
		v[y]*v2[z] - v2[y]*v[z],
		v[z]*v2[x] - v2[z]*v[x],
		v[x]*v2[y] - v2[x]*v[y],
	}, nil
}

// Rotate is rotating a vector around an abitrary vector axis
// If no axis are specified it will default to rotate around the Z axis
//
// If a vector with more than 3-dimensions is rotated, it will cut the extra
// dimensions and return a 3-dimensional vector.
//
// NOTE: the ...Vector is just syntactic sugar that allows the vector axis to not be
// specified and default to the Z axis, if multiple axis is passed the first will be
// set as the rotational axis
func Rotate(v Vector, angle float64, as ...Vector) Vector {
	return v.Clone().Rotate(angle, as...)
}

// Rotate is rotating a vector around an abitrary vector axis
// If no axis are specified it will default to rotate around the Z axis
//
// If a vector with more than 3-dimensions is rotated, it will cut the extra
// dimensions and return a 3-dimensional vector.
//
// NOTE: the ...Vector is just syntactic sugar that allows the vector axis to not be
// specified and default to the Z axis, if multiple axis is passed the first will be
// set as the rotational axis
func (v Vector) Rotate(angle float64, as ...Vector) Vector {
	axis, dim := Z, len(v)

	if len(as) > 0 {
		axis = as[0]
	}

	if dim == 0 {
		return v
	}

	if dim > 3 {
		v = v[:3]
	}

	if l := len(axis); l < 3 {
		axis = append(axis, make([]float64, 3-l)...)
	}

	if dim < 3 {
		v = append(v, make([]float64, 3-dim)...)
	}

	cos, sin, u := math.Cos(angle), math.Sin(angle), Unit(axis)

	x, _ := u.Cross(v)
	d := u.Dot(v)

	v.Sum(v.Scale(cos), x.Scale(sin), u.Scale(d).Scale(1-cos))

	if dim < 3 && axis.Equal(Z) {
		return v[:2]
	}

	return v
}

// Angle returns the angle in radians from the first Vector to the second, and an error if the two Vectors
// aren't of equal dimensions (length). For 0-dimension Vectors, the returned angle is 0. For 1-dimension Vectors,
// the angle is Pi if the second Vector's coordinate is less than the first Vector's coordinate, and 0 otherwise.
func Angle(v1, v2 Vector) (float64, error) {
	return v1.Angle(v2)
}

// Angle returns the angle in radians from the first Vector to the second, and an error if the two Vectors
// aren't of equal dimensions (length). For 0-dimension Vectors, the returned angle is 0. For 1-dimension Vectors,
// the angle is Pi if the second Vector's coordinate is less than the first Vector's coordinate, and 0 otherwise.
func (v Vector) Angle(v2 Vector) (float64, error) {
	dim := len(v)
	dim2 := len(v2)

	if dim != dim2 {
		return 0, ErrNotSameDimensions
	}

	if dim == 0 {
		return 0, nil
	}

	if dim == 1 {
		if v2[x] < v[x] {
			return math.Pi, nil
		}
		return 0, nil
	}

	if dim == 2 {
		return (math.Atan2(v2[y], v2[x]) - math.Atan2(v[y], v[x])), nil
	}

	// 3 or more dimensions
	angle := math.Acos(Dot(Unit(v), Unit(v2)))
	return angle, nil

}

// String returns the string representation of a vector
func (v Vector) String() (str string) {
	if v == nil {
		return "[]"
	}

	for i := range v {
		if v[i] < 1e-8 && v[i] > 0 {
			str += "0 "
		} else {
			str += fmt.Sprint(v[i]) + " "
		}

	}

	return "[" + str[:len(str)-1] + "]"
}

// X is corresponding to doing a v[0] lookup, if index 0 does not exist yet, a
// 0 will be returned instead
func (v Vector) X() float64 {
	if len(v) < 1 {
		return 0.
	}

	return v[x]
}

// Y is corresponding to doing a v[1] lookup, if index 1 does not exist yet, a
// 0 will be returned instead
func (v Vector) Y() float64 {
	if len(v) < 2 {
		return 0.
	}

	return v[y]
}

// Z is corresponding to doing a v[2] lookup, if index 2 does not exist yet, a
// 0 will be returned instead
func (v Vector) Z() float64 {
	if len(v) < 3 {
		return 0.
	}

	return v[z]
}
