package vector_test

import (
	"fmt"
	"testing"

	"github.com/kvartborg/vector"
)

type vec = vector.Vector

// TODO: this test results in a panic on arm64
// func TestSlicingOfVectors(t *testing.T) {
// 	v1 := vec{1, 2, 3}
// 	v2 := v1[1:]
//
// 	result := vector.Add(v1, v2)
//
// 	if result[0] != 3 || result[1] != 5 || result[2] != 3 {
// 		t.Error("vector did not get sliced correctly")
// 	}
// }

func TestMultiDimensionalVec(t *testing.T) {
	result := vec{1}.Add(vec{1, 2})

	if len(result) > 1 {
		t.Error("did not normalise vector to lowest dimension")
	}
}

func Example() {
	// Create a new vector
	v1 := vec{4, 2}

	// Create a vector from a list of float64
	v2 := vec([]float64{1, 2, 4})

	fmt.Println(
		v1.Add(v2),
	)
	// Output: [5 4]
}

func ExampleAdd() {
	fmt.Println(
		vector.Add(vec{0, 2}, vec{1, 4}),
	)
	// Output: [1 6]
}

func ExampleVector_Add() {
	fmt.Println(
		vec{0, 2}.Add(vec{1, 4}),
	)
	// Output: [1 6]
}

func ExampleSub() {
	fmt.Println(
		vector.Sub(vec{1, 4}, vec{0, 2}),
	)
	// Output: [1 2]
}

func ExampleVector_Sub() {
	fmt.Println(
		vec{1, 4}.Sub(vec{0, 2}),
	)
	// Output: [1 2]
}

func ExampleEqual() {
	fmt.Println(
		vec{1, 2}.Equal(vec{1, 2}),
	)
	// Output: true
}

func ExampleVector_Equal() {
	fmt.Println(
		vec{2, 1}.Equal(vec{1, 2}),
	)
	// Output: false
}

func ExampleDot() {
	fmt.Println(
		vector.Dot(vec{0, 2}, vec{2, 0}),
	)
	// Output: 0
}

func ExampleVector_Dot() {
	fmt.Println(
		vec{0, 2}.Dot(vec{2, 0}),
	)
	// Output: 0
}

func ExampleCross() {
	fmt.Println(
		vector.Cross(vec{0, 1, 2}, vec{3, 2, 1}),
	)
	// Output: [-3 6 -6] <nil>
}

func ExampleVector_Cross() {
	fmt.Println(
		vec{0, 1, 2}.Cross(vec{3, 2, 1}),
	)
	// Output: [-3 6 -6] <nil>
}

func ExampleClone() {
	fmt.Println(
		vector.Clone(vec{1, 2}),
	)
	// Output: [1 2]
}

func ExampleVector_Clone() {
	fmt.Println(
		vec{1, 2}.Clone(),
	)
	// Output: [1 2]
}

func ExampleScale() {
	fmt.Println(
		vector.Scale(vec{1, 2}, 2),
	)
	// Output: [2 4]
}

func ExampleVector_Scale() {
	fmt.Println(
		vec{1, 2}.Scale(2),
	)
	// Output: [2 4]
}

func ExampleMagnitude() {
	fmt.Println(
		vector.Magnitude(vec{1, 2}),
	)
	// Output: 2.23606797749979
}

func ExampleVector_Magnitude() {
	fmt.Println(
		vec{1, 2}.Magnitude(),
	)
	// Output: 2.23606797749979
}

func ExampleUnit() {
	fmt.Println(
		vector.Unit(vec{1, 2}),
	)
	// Output: [0.4472135954999579 0.8944271909999159]
}

func ExampleVector_Unit() {
	fmt.Println(
		vec{1, 2}.Unit(),
	)
	// Output: [0.4472135954999579 0.8944271909999159]
}

func BenchmarkAdd(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Add(v1, v2)
	}
}

func BenchmarkVector_Add(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}

func BenchmarkSub(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Sub(v2, v1)
	}
}

func BenchmarkVector_Sub(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		v2.Sub(v1)
	}
}

func BenchmarkClone(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Clone(v)
	}
}

func BenchmarkVector_Clone(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Clone()
	}
}

func BenchmarkScale(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Scale(v, 2)
	}
}

func BenchmarkVector_Scale(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Scale(2)
	}
}

func BenchmarkEqual(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Equal(v1, v2)
	}
}

func BenchmarkMagnitude(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Magnitude(v)
	}
}

func BenchmarkDot(b *testing.B) {
	v1, v2 := vec{1, 2}, vec{2, 1}

	for i := 0; i < b.N; i++ {
		vector.Dot(v1, v2)
	}
}

func BenchmarkCross(b *testing.B) {
	v1, v2 := vec{1, 2, 3}, vec{3, 2, 1}

	for i := 0; i < b.N; i++ {
		vector.Cross(v1, v2)
	}
}

func BenchmarkUnit(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Unit(v)
	}
}

func BenchmarkVector_Unit(b *testing.B) {
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Unit()
	}
}

func BenchmarkStructVectorAddition(b *testing.B) {
	type vector struct{ x, y, z float64 }

	add := func(v1, v2 vector) vector {
		return vector{
			v1.x + v2.x,
			v1.y + v2.y,
			v1.z + v2.z,
		}
	}

	v1, v2 := vector{}, vector{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(v1, v2)
	}
}
