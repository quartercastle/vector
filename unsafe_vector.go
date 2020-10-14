package vector

type UnsafeVector []float64

func In(a Vector) UnsafeVector {
	return UnsafeVector(a)
}

func (a UnsafeVector) Clone() UnsafeVector {
	return clone(a)
}

func (a UnsafeVector) Add(b Vector) UnsafeVector {
	return add(a, b)
}

// Sum a vector with a vector or a set of vectors
func (a UnsafeVector) Sum(vectors ...Vector) UnsafeVector {
	return sum(a, vectors)
}

// Sub subtracts a vector with another vector or a set of vectors
func (a UnsafeVector) Sub(b Vector) UnsafeVector {
	return sub(a, b)
}

// Invert inverts the vector, and then returns it
func (a UnsafeVector) Invert() UnsafeVector {
	return invert(a)
}

// Scale vector with a given size
func (a UnsafeVector) Scale(size float64) UnsafeVector {
	return scale(a, size)
}

// Equal compares that two vectors are equal to each other
func (a UnsafeVector) Equal(b Vector) bool {
	return equal(a, b)
}

// Magnitude of a vector
func (a UnsafeVector) Magnitude() float64 {
	return magnitude(a)
}

// Unit returns a direction vector with the length of one.
func (a UnsafeVector) Unit() UnsafeVector {
	return unit(a)
}

// Dot product of two vectors
func (a UnsafeVector) Dot(b Vector) float64 {
	return dot(a, b)
}

// Cross product of two vectors
func (a UnsafeVector) Cross(b Vector) (Vector, error) {
	return cross(a, b)
}

// Rotate is rotating a vector around an abitrary vector axis
// If no axis are specified it will default to rotate around the Z axis
//
// If a vector with more than 3-dimensions is rotated, it will cut the extra
// dimensions and return a 3-dimensional vector.
//
// NOTE: the ...UnsafeVector is just syntactic sugar that allows the vector axis to not be
// specified and default to the Z axis, if multiple axis is passed the first will be
// set as the rotational axis
func (a UnsafeVector) Rotate(angle float64, axis ...Vector) UnsafeVector {
	as := Z

	if len(axis) > 0 {
		as = axis[0]
	}

	return rotate(a, angle, clone(as))
}

// Angle returns the angle in radians from the first UnsafeVector to the second, and an error if the two UnsafeVectors
// aren't of equal dimensions (length). For 0-dimension UnsafeVectors, the returned angle is 0. For 1-dimension UnsafeVectors,
// the angle is Pi if the second UnsafeVector's coordinate is less than the first UnsafeVector's coordinate, and 0 otherwise.
func (a UnsafeVector) Angle(axis ...Vector) float64 {
	as := Z

	if len(axis) > 0 {
		as = axis[0]
	}

	return angle(a, as)
}

// X is corresponding to doing a UnsafeVector[0] lookup, if index 0 does not exist yet, a
// 0 will be returned instead
func (a UnsafeVector) X() float64 {
	if len(a) < 1 {
		return 0
	}

	return a[x]
}

// Y is corresponding to doing a UnsafeVector[1] lookup, if index 1 does not exist yet, a
// 0 will be returned instead
func (a UnsafeVector) Y() float64 {
	if len(a) < 2 {
		return 0
	}

	return a[y]
}

// Z is corresponding to doing a UnsafeVector[2] lookup, if index 2 does not exist yet, a
// 0 will be returned instead
func (a UnsafeVector) Z() float64 {
	if len(a) < 3 {
		return 0
	}

	return a[z]
}
