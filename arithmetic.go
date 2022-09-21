package vector

import (
	"math"
)

const (
	x = iota
	y
	z
)

func clone(a []float64) []float64 {
	clone := make([]float64, len(a))
	copy(clone, a)
	return clone
}

func add(a, b []float64) []float64 {
	dimA, dimB := len(a), len(b)

	if (dimA == 1 || dimA == 2 || dimA == 3) && dimB == 1 {
		a[x] += b[x]
		return a
	}

	if dimA == 2 && dimB == 2 {
		a[x], a[y] = a[x]+b[x], a[y]+b[y]
		return a
	}

	if dimA == 3 && dimB == 2 {
		a[x], a[y] = a[x]+b[x], a[y]+b[y]
		return a
	}

	if dimA == 3 && dimB == 3 {
		a[x], a[y], a[z] = a[x]+b[x], a[y]+b[y], a[z]+b[z]
		return a
	}

	if dimB > dimA {
		axpyUnitaryTo(a, 1, a, b[:dimA])
	} else {
		axpyUnitaryTo(a, 1, a, b)
	}

	return a
}

func sum(a []float64, vectors []Vector) []float64 {
	dim := len(a)

	if (dim == 1 || dim == 2 || dim == 3) && len(vectors) == 1 && len(vectors[0]) == 1 {
		a[x] += vectors[0][x]
		return a
	}

	if dim == 2 && len(vectors) == 1 && len(vectors[0]) == 2 {
		a[x], a[y] = a[x]+vectors[0][x], a[y]+vectors[0][y]
		return a
	}

	if dim == 3 && len(vectors) == 1 && len(vectors[0]) == 2 {
		a[x], a[y] = a[x]+vectors[0][x], a[y]+vectors[0][y]
		return a
	}

	if dim == 3 && len(vectors) == 1 && len(vectors[0]) == 3 {
		a[x], a[y], a[z] = a[x]+vectors[0][x], a[y]+vectors[0][y], a[z]+vectors[0][z]
		return a
	}

	for i := range vectors {
		if len(vectors[i]) > dim {
			axpyUnitaryTo(a, 1, a, vectors[i][:dim])
		} else {
			axpyUnitaryTo(a, 1, a, vectors[i])
		}
	}

	return a
}

func sub(a, b []float64) []float64 {
	dimA, dimB := len(a), len(b)

	if (dimA == 1 || dimA == 2 || dimA == 3) && dimB == 1 {
		a[x] -= b[x]
		return a
	}

	if dimA == 2 && dimB == 2 {
		a[x], a[y] = a[x]-b[x], a[y]-b[y]
		return a
	}

	if dimA == 3 && dimB == 1 {
		a[x] -= b[x]
		return a
	}

	if dimA == 3 && dimB == 2 {
		a[x], a[y] = a[x]-b[x], a[y]-b[y]
		return a
	}

	if dimA == 3 && dimB == 3 {
		a[x], a[y], a[z] = a[x]-b[x], a[y]-b[y], a[z]-b[z]
		return a
	}

	if dimB > dimA {
		axpyUnitaryTo(a, -1, b[:dimA], a)
	} else {
		axpyUnitaryTo(a, -1, b, a)
	}

	return a
}

func invert(a []float64) []float64 {
	for i := range a {
		a[i] *= -1
	}
	return a
}

func scale(a []float64, size float64) []float64 {
	dim := len(a)

	if dim == 2 {
		a[x], a[y] = a[x]*size, a[y]*size
		return a
	}

	if dim == 3 {
		a[x], a[y], a[z] = a[x]*size, a[y]*size, a[z]*size
		return a
	}

	scalUnitaryTo(a, size, a)
	return a
}

func equal(a, b []float64) bool {
	dim := len(a)
	if dim != len(b) {
		return false
	}

	if dim == 2 {
		return math.Abs(a[x]-b[x]) < 1e-8 && math.Abs(a[y]-b[y]) < 1e-8
	}

	if dim == 3 {
		return math.Abs(a[x]-b[x]) < 1e-8 && math.Abs(a[y]-b[y]) < 1e-8 && math.Abs(a[z]-b[z]) < 1e-8
	}

	for i := range a {
		if math.Abs(a[i]-b[i]) > 1e-8 {
			return false
		}
	}

	return true
}

func magnitude(a []float64) float64 {
	dim := len(a)

	if dim == 1 {
		return math.Sqrt(a[x] * a[x])
	}

	if dim == 2 {
		return math.Sqrt(a[x]*a[x] + a[y]*a[y])
	}

	if dim == 3 {
		return math.Sqrt(a[x]*a[x] + a[y]*a[y] + a[z]*a[z])
	}

	var result float64
	for _, scalar := range a {
		result += scalar * scalar
	}

	return math.Sqrt(result)
}

func unit(a []float64) []float64 {
	dim := len(a)

	if dim == 2 {
		l := math.Sqrt(a[x]*a[x] + a[y]*a[y])
		if l < 1e-8 {
			return a
		}
		a[x], a[y] = a[x]/l, a[y]/l
		return a
	}

	if dim == 3 {
		l := math.Sqrt(a[x]*a[x] + a[y]*a[y] + a[z]*a[z])
		if l < 1e-8 {
			return a
		}
		a[x], a[y], a[z] = a[x]/l, a[y]/l, a[z]/l
		return a
	}

	l := magnitude(a)

	if l < 1e-8 {
		return a
	}

	for i := range a {
		a[i] = a[i] / l
	}

	return a
}

func dot(a, b []float64) float64 {
	result, dimA, dimB := 0., len(a), len(b)

	if dimA == 2 && dimB == 2 {
		result = a[x]*b[x] + a[y]*b[y]
		if result > 1 {
			result = 1
		} else if result < -1 {
			result = -1
		}
		return result
	}

	if dimA == 3 && dimB == 3 {
		result = a[x]*b[x] + a[y]*b[y] + a[z]*b[z]
		if result > 1 {
			result = 1
		} else if result < -1 {
			result = -1
		}
		return result
	}

	if dimA > dimB {
		b = append(b, make(Vector, dimA-dimB)...)
	}

	if dimA < dimB {
		a = append(a, make(Vector, dimB-dimA)...)
	}

	for i := range a {
		result += a[i] * b[i]
	}

	if result > 1 {
		result = 1
	} else if result < -1 {
		result = -1
	}

	return result
}

func cross(a, b []float64) ([]float64, error) {
	if len(a) != 3 || len(b) != 3 {
		return nil, ErrNot3Dimensional
	}

	return []float64{
		a[y]*b[z] - b[y]*a[z],
		a[z]*b[x] - b[z]*a[x],
		a[x]*b[y] - b[x]*a[y],
	}, nil
}

func rotate(a []float64, angle float64, axis []float64) []float64 {
	dim := len(a)

	if dim == 0 {
		return a
	}

	cos, sin := math.Cos(angle), math.Sin(angle)

	if dim == 1 && equal(axis, Z) {
		a = append(a, 0)
	}

	if (dim == 1 || dim == 2) && equal(axis, Z) {
		ax, ay := a[x], a[y]
		a[x] = ax*cos - ay*sin
		a[y] = ax*sin + ay*cos
		return a
	}

	if dim == 3 && equal(axis, X) {
		ay, az := a[y], a[z]
		a[y] = ay*cos - az*sin
		a[z] = ay*sin + az*cos
		return a
	}

	if dim == 3 && equal(axis, Y) {
		ax, az := a[x], a[z]
		a[x] = ax*cos + az*sin
		a[z] = -ax*sin + az*cos
		return a
	}

	if dim == 3 && equal(axis, Z) {
		ax, ay := a[x], a[y]
		a[x] = ax*cos - ay*sin
		a[y] = ax*sin + ay*cos
		return a
	}

	if dim > 3 {
		a = a[:3]
	}

	if l := len(axis); l < 3 {
		axis = append(axis, make([]float64, 3-l)...)
	}

	if dim < 3 {
		a = append(a, make([]float64, 3-dim)...)
	}

	u := unit(clone(axis))

	x, _ := cross(u, a)
	d := dot(u, a)

	add(a, scale(a, cos))
	add(a, scale(x, sin))
	add(a, scale(scale(u, d), 1-cos))

	if dim < 3 && equal(axis, Z) {
		return a[:2]
	}

	return a
}

func angle(a, b []float64) float64 {
	dimA, dimB := len(a), len(b)

	if dimA == 0 {
		return 0
	}

	if dimA == 1 && dimB == 1 {
		if b[x] < a[x] {
			return math.Pi
		}
		return 0
	}

	if dimA == 2 && equal(b, X) {
		return math.Atan2(a[y], a[x])
	}

	if dimA < dimB {
		a = append(a, make([]float64, dimB-dimA)...)
	}

	if dimA > dimB {
		b = append(b, make([]float64, dimA-dimB)...)
	}

	if dimA == 2 {
		return (math.Atan2(b[y], b[x]) - math.Atan2(a[y], a[x]))
	}

	// 3 or more dimensions
	angle := math.Acos(dot(unit(clone(a)), unit(clone(b))))
	return angle
}

func swizzle(a []float64, indices ...int) ([]float64, error) {

	for _, i := range indices {
		if i < 0 || i >= len(a) {
			return nil, ErrNotValidSwizzleIndex
		}
	}

	switch len(indices) {

	case 0:
		return []float64{}, nil
	case 1:
		return []float64{a[indices[0]]}, nil
	case 2:
		return []float64{a[indices[0]], a[indices[1]]}, nil
	case 3:
		return []float64{a[indices[0]], a[indices[1]], a[indices[2]]}, nil
	case 4:
		return []float64{a[indices[0]], a[indices[1]], a[indices[2]], a[indices[3]]}, nil
	default:

		vec := make([]float64, len(indices))

		for i := 0; i < len(indices); i++ {

			ind := indices[i]
			if ind >= len(a) {
				return nil, ErrNotValidSwizzleIndex
			}
			vec[i] = a[ind]

		}

		return vec, nil

	}

}
