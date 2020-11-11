package vector_test

import (
	"math"
	"testing"

	"github.com/kvartborg/vector"
)

type vec = vector.Vector

func TestCasting(t *testing.T) {
	result := vec{1, 2}.Sum([]float64{2, 4})

	if !result.Equal(vec{3, 6}) {
		t.Errorf("casting did not work as expected")
	}
}

func TestUnsafeVector(t *testing.T) {
	result := make(vec, 2)
	vector.In(result).Sum(vec{1, 2}, vec{1, 2})
	if result.X() != 2 || result.Y() != 4 {
		t.Error("UnsafeVector did not work as expected")
	}
}

func TestSlicingOfVectors(t *testing.T) {
	v1 := vec{1, 2, 3}
	v2 := v1[1:]

	result := v1.Add(v2)

	if result[0] != 3 || result[1] != 5 || result[2] != 3 {
		t.Error("vector did not get sliced correctly")
	}
}

func TestMultiDimensionalVec(t *testing.T) {
	v1 := vec{1}.Sum(vec{1, 2})
	v2 := vec{1, 2}.Sum(vec{1})

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

	if result.X() > 1e-8 || result.Y() != 0 || result.Z() != -1 {
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
