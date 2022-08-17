package vector

// Vector is the definition of a row vector that contains scalars as
// 64 bit floats
type Vector []float64

// Clone a vector
func (a Vector) Clone() Vector {
	return clone(a)
}

func (a Vector) Add(b Vector) Vector {
	return add(clone(a), b)
}

// Sum a vector with a vector or a set of vectors
func (a Vector) Sum(vectors ...Vector) Vector {
	return sum(clone(a), vectors)
}

// Sub subtracts a vector with another vector or a set of vectors
func (a Vector) Sub(b Vector) Vector {
	return sub(clone(a), b)
}

// Invert inverts the vector, and then returns it
func (a Vector) Invert() Vector {
	return invert(clone(a))
}

// Scale vector with a given size
func (a Vector) Scale(size float64) Vector {
	return scale(clone(a), size)
}

// Equal compares that two vectors are equal to each other
func (a Vector) Equal(b Vector) bool {
	return equal(a, b)
}

// Magnitude of a vector
func (a Vector) Magnitude() float64 {
	return magnitude(a)
}

// Unit returns a direction vector with the length of one.
func (a Vector) Unit() Vector {
	return unit(clone(a))
}

// Dot product of two vectors
func (a Vector) Dot(b Vector) float64 {
	return dot(a, b)
}

// Cross product of two vectors
func (a Vector) Cross(b Vector) (Vector, error) {
	return cross(a, b)
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
func (a Vector) Rotate(angle float64, axis ...Vector) Vector {
	as := Z

	if len(axis) > 0 {
		as = axis[0]
	}

	return rotate(clone(a), angle, clone(as))
}

// Swizzle returns a clone of the input vector altered using the provided swizzling indices.
// For example, with `vector := {1, 3, 9}`, `vector.Swizzle(2,1,2,0)` will return `Vector{9,3,9,1}`.
// Swizzle will return the swizzled vector, and an error if one of the provided indices is out of bounds.
func (a Vector) Swizzle(swizzleIndices ...int) (Vector, error) {
	return swizzle(a, swizzleIndices...)
}

// Angle returns the angle in radians from the first Vector to the second, and an error if the two Vectors
// aren't of equal dimensions (length). For 0-dimension Vectors, the returned angle is 0. For 1-dimension Vectors,
// the angle is Pi if the second Vector's coordinate is less than the first Vector's coordinate, and 0 otherwise.
func (a Vector) Angle(axis ...Vector) float64 {
	as := X

	if len(axis) > 0 {
		as = axis[0]
	}

	return angle(a, as)
}

// X is corresponding to doing a Vector[0] lookup, if index 0 does not exist yet, a
// 0 will be returned instead
func (a Vector) X() float64 {
	if len(a) < 1 {
		return 0
	}

	return a[x]
}

// Y is corresponding to doing a Vector[1] lookup, if index 1 does not exist yet, a
// 0 will be returned instead
func (a Vector) Y() float64 {
	if len(a) < 2 {
		return 0
	}

	return a[y]
}

// Z is corresponding to doing a Vector[2] lookup, if index 2 does not exist yet, a
// 0 will be returned instead
func (a Vector) Z() float64 {
	if len(a) < 3 {
		return 0
	}

	return a[z]
}
