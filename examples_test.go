package vector_test

import (
	"fmt"
	"math"

	"github.com/kvartborg/vector"
)

func Example() {
	// create a zero vector of 3-dimensions
	v1 := make(vec, 3)

	// Create a new vector
	v2 := vec{4, 2}

	// Create a vector from a list of float64
	v3 := vec([]float64{1, 2, 4})

	fmt.Println(
		v1.Sum(v2, v3),
	)
	// Output: [5 4 4]
}

func ExampleInPlace() {
	result, v1, v2 := make(vec, 3), vec{2, 1, 3}, vec{4, 12, 6}

	vector.In(result).
		Add(v2).
		Sub(v1)

	fmt.Println(result)
	// Output: [2 11 3]
}

func ExampleVector_Sum() {
	fmt.Println(
		vec{0, 2}.Sum(vec{1, 4}),
	)
	// Output: [1 6]
}

func ExampleVector_Sub() {
	fmt.Println(
		vec{1, 4}.Sub(vec{0, 2}),
	)
	// Output: [1 2]
}

func ExampleVector_Equal() {
	fmt.Println(
		vec{2, 1}.Equal(vec{1, 2}),
	)
	// Output: false
}

func ExampleVector_Dot() {
	fmt.Println(
		vec{0, 2}.Dot(vec{2, 0}),
	)
	// Output: 0
}

func ExampleVector_Cross() {
	fmt.Println(
		vec{0, 1, 2}.Cross(vec{3, 2, 1}),
	)
	// Output: [-3 6 -3] <nil>
}

func ExampleVector_Clone() {
	fmt.Println(
		vec{1, 2}.Clone(),
	)
	// Output: [1 2]
}

func ExampleVector_Scale() {
	fmt.Println(
		vec{1, 2}.Scale(2),
	)
	// Output: [2 4]
}

func ExampleVector_Magnitude() {
	fmt.Println(
		vec{1, 2}.Magnitude(),
	)
	// Output: 2.23606797749979
}

func ExampleVector_Unit() {
	fmt.Println(
		vec{1, 2}.Unit(),
	)
	// Output: [0.4472135954999579 0.8944271909999159]
}

func ExampleVector_Rotate() {
	fmt.Println(
		vec{1, 0, 0}.Rotate(math.Pi/2, vector.Y),
	)
	// Output: [6.123233995736757e-17 0 -1]
}

func ExampleVector_Angle() {
	fmt.Println(
		vec{17, 4, 3}.Angle(vec{-1, 15, 7}),
	)
	// Output: 1.351241200672429
}

func ExampleVector_Invert() {
	fmt.Println(
		vec{19, 0, 3}.Invert(),
	)
	// Output: [-19 -0 -3]
}
