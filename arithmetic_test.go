package vector

import (
	"math"
	"testing"
)

func BenchmarkArithmeticAdd(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := []float64{1, 2}, []float64{2, 3}

	for i := 0; i < b.N; i++ {
		add(v1, v2)
	}
}

func BenchmarkArithmeticSum(b *testing.B) {
	b.ReportAllocs()
	v1, vectors := []float64{1, 2}, []Vector{{2, 3}}

	for i := 0; i < b.N; i++ {
		sum(v1, vectors)
	}
}

func BenchmarkArithmeticSub(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := []float64{1, 2}, []float64{2, 3}

	for i := 0; i < b.N; i++ {
		sub(v2, v1)
	}
}

func BenchmarkArithmeticScale(b *testing.B) {
	b.ReportAllocs()
	v := []float64{1, 2}

	for i := 0; i < b.N; i++ {
		scale(v, 2)
	}
}

func BenchmarkArithmeticEqual(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := []float64{1, 2}, []float64{1, 2}

	for i := 0; i < b.N; i++ {
		equal(v1, v2)
	}
}

func BenchmarkArithmeticMagnitude(b *testing.B) {
	b.ReportAllocs()
	v := []float64{1, 2}

	for i := 0; i < b.N; i++ {
		magnitude(v)
	}
}

func BenchmarkArithmeticDot(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := []float64{1, 2}, []float64{2, 1}

	for i := 0; i < b.N; i++ {
		dot(v1, v2)
	}
}

func BenchmarkArithmeticCross(b *testing.B) {
	b.ReportAllocs()
	v1, v2 := []float64{1, 2, 3}, []float64{3, 2, 1}

	for i := 0; i < b.N; i++ {
		cross(v1, v2)
	}
}

func BenchmarkArithmeticUnit(b *testing.B) {
	b.ReportAllocs()
	v := []float64{1, 2}

	for i := 0; i < b.N; i++ {
		unit(v)
	}
}

func BenchmarkArithmeticRotate(b *testing.B) {
	b.ReportAllocs()
	v := []float64{1, 2}

	for i := 0; i < b.N; i++ {
		rotate(v, math.Pi/2, []float64{0, 0, 1})
	}
}

func BenchmarkArithmeticAngle(b *testing.B) {
	b.ReportAllocs()

	v1 := []float64{0, 4}
	v2 := []float64{1, 0, 0}

	for i := 0; i < b.N; i++ {
		angle(v1, v2)
	}
}

func BenchmarkSwizzling(b *testing.B) {
	b.ReportAllocs()

	v1 := Vector{1, 2, 19}

	for i := 0; i < b.N; i++ {
		swizzle(v1, 0, 1, 2)
	}

}
