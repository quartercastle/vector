package vector_test

import (
	"math"
	"testing"

	"github.com/kvartborg/vector"
)

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

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		vector.Sum(v1, v2)
	}
}

func BenchmarkVector_Sum(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := vec{1, 2}, vec{2, 3}

	for i := 0; i < b.N; i++ {
		v1.Sum(v2)
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

func BenchmarkVectorAngle(b *testing.B) {
	b.ReportAllocs()
	v := vec{0, 1}
	v2 := vec{0.7, 0.3}

	for i := 0; i < b.N; i++ {
		vector.Angle(v, v2)
	}
}

func BenchmarkVector_Angle(b *testing.B) {
	b.ReportAllocs()

	v := vec{0, 4, 3}
	v2 := vec{0, -4, -3}

	for i := 0; i < b.N; i++ {
		v.Angle(v2)
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

	v1, v2 := vector{1, 2, 3}, vector{1, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(v1, v2)
	}
}

func BenchmarkListVectorAddition(b *testing.B) {
	b.ReportAllocs()
	type vector []float64

	add := func(v1, v2 vector) vector {
		return vector{
			v1[0] + v2[0],
			v1[1] + v2[1],
			v1[2] + v2[2],
		}
	}

	v1, v2 := vector{1, 2, 3}, vector{1, 2, 3}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(v1, v2)
	}
}
