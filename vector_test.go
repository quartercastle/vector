package vector_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/kvartborg/vector"
)

type vec = vector.Vector

func TestCasting(t *testing.T) {
	result := vec{1, 2}.Add([]float64{2, 4})

	if !result.Equal(vec{3, 6}) {
		t.Errorf("Casting did not work as expected")
	}
}

func TestSlicingOfVectors(t *testing.T) {
	v1 := vec{1, 2, 3}
	v2 := v1[1:]

	result := vector.Add(v1, v2)

	if result[0] != 3 || result[1] != 5 || result[2] != 3 {
		t.Error("vector did not get sliced correctly")
	}
}

func TestMultiDimensionalVec(t *testing.T) {
	v1 := vec{1}.Add(vec{1, 2})
	v2 := vec{1, 2}.Add(vec{1})

	if len(v1) != 1 || len(v2) != 2 {
		t.Error("did not normalise vector to lowest dimension")
	}
}

func TestRotationOfVector(t *testing.T) {
	result := vec{1}.Rotate(math.Pi / 2)

	if result.X() > 1e-8 || result.Y() != 1 {
		t.Error("did not up scale to 2-dimensions")
	}

	result = vec{1}.Rotate(math.Pi/2, vector.Y)

	if result.X() > 1e-8 || result.Y() > 1e-8 || result.Z() != -1 {
		t.Error("did not up scale to 2-dimensions")
	}

	result = vec{1, 0}.Rotate(math.Pi / 2)

	if result.X() > 1e-8 || result.Y() != 1 || len(result) != 2 {
		t.Error("did not keep 2-dimensions when input vec is 2-dimensions and it should rotate around the z axis")
	}

	result = vec{1, 0}.Rotate(math.Pi/2, vector.Y)

	if result.X() > 1e-8 || result.Y() != 0 || result[vector.Z] != -1 {
		t.Error("did not upscale to 3-dimensions")
	}

	result = vec{1, 0, 0, 0}.Rotate(math.Pi/2, vector.Y)

	if len(result) > 3 {
		t.Error("did not cut extra dimensions")
	}

}

func TestXYZGetters(t *testing.T) {
	v1 := vec{}

	if v1.X() != 0 || v1.Y() != 0 || v1.Z() != 0 {
		t.Error("getter methods for x, y, z did not return 0 when expected")
	}

	v2 := vec{1, 2, 3}

	if v2.X() != 1 || v2.Y() != 2 || v2.Z() != 3 {
		t.Error("getter methods for x, y, z did not return 0 when expected")
	}
}

func Example() {
	// create a zero vector of 3-dimensions
	v1 := make(vec, 3)

	// Create a new vector
	v2 := vec{4, 2}

	// Create a vector from a list of float64
	v3 := vec([]float64{1, 2, 4})

	fmt.Println(
		v1.Add(v2, v3),
	)
	// Output: [5 4 4]
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
	// Output: [-3 6 -3] <nil>
}

func ExampleVector_Cross() {
	fmt.Println(
		vec{0, 1, 2}.Cross(vec{3, 2, 1}),
	)
	// Output: [-3 6 -3] <nil>
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

func ExampleRotate() {
	fmt.Println(
		vector.Rotate(vec{1, 0}, math.Pi/2),
	)
	// Output: [0 1]
}

func ExampleVector_Rotate() {
	fmt.Println(
		vec{1, 0, 0}.Rotate(math.Pi/2, vector.Y),
	)
	// Output: [0 0 -1]
}

func ExampleAngle() {
	fmt.Println(
		vector.Angle(vec{1, 0}, vec{0, 1}),
	)
	// Output: 1.5707963267948966 [0 0 1] <nil>
}

func ExampleVector_Angle() {
	fmt.Println(
		vec{17, 4, 3}.Angle(vec{-1, 15, 7}),
	)
	// Output: 1.351241200672429 [-0.05927484240309407 -0.42538416312808686 0.9030696577883156] <nil>
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Add(v1, v2)
	}
}

func BenchmarkVector_Add(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		v1.Add(v2)
	}
}

func BenchmarkSub(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Sub(v2, v1)
	}
}

func BenchmarkVector_Sub(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		v2.Sub(v1)
	}
}

func BenchmarkClone(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Clone(v)
	}
}

func BenchmarkVector_Clone(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Clone()
	}
}

func BenchmarkScale(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Scale(v, 2)
	}
}

func BenchmarkVector_Scale(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Scale(2)
	}
}

func BenchmarkEqual(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Equal(v1, v2)
	}
}

func BenchmarkMagnitude(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Magnitude(v)
	}
}

func BenchmarkDot(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 1}

	for i := 0; i < b.N; i++ {
		vector.Dot(v1, v2)
	}
}

func BenchmarkCross(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2, 3}, vec{3, 2, 1}

	for i := 0; i < b.N; i++ {
		vector.Cross(v1, v2)
	}
}

func BenchmarkUnit(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Unit(v)
	}
}

func BenchmarkVector_Unit(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Unit()
	}
}

func BenchmarkVectorRotate(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		vector.Rotate(v, math.Pi/2)
	}
}

func BenchmarkVector_Rotate(b *testing.B) {
	b.ReportAllocs()
	v := vec{1, 2}

	for i := 0; i < b.N; i++ {
		v.Rotate(math.Pi / 2)
	}
}

func BenchmarkSlicing(b *testing.B) {
	b.ReportAllocs()
	v := make(vec, 100)

	for i := 0; i < b.N; i++ {
		_ = v[:]
	}
}

func BenchmarkStructVectorAddition(b *testing.B) {
	b.ReportAllocs()
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
