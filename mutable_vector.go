package vector

// MutableVector is a vector where all arithmetic operations will be done in
// place on the calling vector. This will increase performance and minimize the
// memory consumption.
type MutableVector []float64

// In takes a vector and turns it into a mutable vector.
func In(a Vector) MutableVector {
	return MutableVector(a)
}

// Clone a mutable vector.
func (a MutableVector) Clone() MutableVector {
	return clone(a)
}

// Add a vector in place in the mutable vector
func (a MutableVector) Add(b Vector) MutableVector {
	return add(a, b)
}

// Sum a vector with a vector or a set of vectors
func (a MutableVector) Sum(vectors ...Vector) MutableVector {
	return sum(a, vectors)
}

// Sub subtracts a vector with another vector or a set of vectors
func (a MutableVector) Sub(b Vector) MutableVector {
	return sub(a, b)
}

// Invert inverts the vector, and then returns it
func (a MutableVector) Invert() MutableVector {
	return invert(a)
}

// Scale vector with a given size
func (a MutableVector) Scale(size float64) MutableVector {
	return scale(a, size)
}

// Equal compares that two vectors are equal to each other
func (a MutableVector) Equal(b Vector) bool {
	return equal(a, b)
}

// Magnitude of a vector
func (a MutableVector) Magnitude() float64 {
	return magnitude(a)
}

// Unit returns a direction vector with the length of one.
func (a MutableVector) Unit() MutableVector {
	return unit(a)
}

// Dot product of two vectors
func (a MutableVector) Dot(b Vector) float64 {
	return dot(a, b)
}

// Cross product of two vectors
func (a MutableVector) Cross(b Vector) (Vector, error) {
	return cross(a, b)
}

// Rotate is rotating a vector around an abitrary vector axis
// If no axis are specified it will default to rotate around the Z axis
//
// If a vector with more than 3-dimensions is rotated, it will cut the extra
// dimensions and return a 3-dimensional vector.
//
// NOTE: the ...MutableVector is just syntactic sugar that allows the vector axis to not be
// specified and default to the Z axis, if multiple axis is passed the first will be
// set as the rotational axis
func (a MutableVector) Rotate(angle float64, axis ...Vector) MutableVector {
	as := Z

	if len(axis) > 0 {
		as = axis[0]
	}

	return rotate(a, angle, clone(as))
}

// Angle returns the angle in radians from the first MutableVector to the second, and an error if the two MutableVectors
// aren't of equal dimensions (length). For 0-dimension MutableVectors, the returned angle is 0. For 1-dimension MutableVectors,
// the angle is Pi if the second MutableVector's coordinate is less than the first MutableVector's coordinate, and 0 otherwise.
func (a MutableVector) Angle(axis ...Vector) float64 {
	as := X

	if len(axis) > 0 {
		as = axis[0]
	}

	return angle(a, as)
}

// X is corresponding to doing a MutableVector[0] lookup, if index 0 does not exist yet, a
// 0 will be returned instead
func (a MutableVector) X() float64 {
	if len(a) < 1 {
		return 0
	}

	return a[x]
}

// Y is corresponding to doing a MutableVector[1] lookup, if index 1 does not exist yet, a
// 0 will be returned instead
func (a MutableVector) Y() float64 {
	if len(a) < 2 {
		return 0
	}

	return a[y]
}

// Z is corresponding to doing a MutableVector[2] lookup, if index 2 does not exist yet, a
// 0 will be returned instead
func (a MutableVector) Z() float64 {
	if len(a) < 3 {
		return 0
	}

	return a[z]
}
