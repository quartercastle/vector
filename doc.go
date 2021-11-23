// Package vector provides useful math operations for vectors and a way to
// represent vectors as a list of float64 values.
//
//		// Minimize the verbosity by using type aliasing
//		type vec = vector.Vector
//
//		// Create a vector
//		v1 := vec{1, 2}
//
//		// Create a vector from a list of float64 values
//		v2 := vec([]float64{2, 6})
//
//		// Do arithmetic operations with the vectors
//		result := v1.Add(v2)
package vector
