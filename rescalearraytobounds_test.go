package rescale

import (
	"testing"
	"reflect"
)

func TestRescaleBoundArray_Shrink(t *testing.T) {
	old_points  := []float64{0.0,25.0,50.0,100.0}
	new_points := RescaleArrayToBounds(0.0, 10.0, old_points)
	expected_new_points  := []float64{0.0,2.5,5.0,10.0}
	
	if (!reflect.DeepEqual(new_points, expected_new_points)) {
		t.Errorf("New point was incorrect, got: %f, want: %f.", new_points, expected_new_points)
	}
}

func TestRescaleBoundArray_Rebase(t *testing.T) {
	old_points  := []float64{-10.0, -9.0, -5.0, 2.0, 6.0}
	new_points := RescaleArrayToBounds(0.0, 100.0, old_points)
	expected_new_points  := []float64{0.00, 6.25, 31.25, 75.00, 100.00}
	
	if (!reflect.DeepEqual(new_points, expected_new_points)) {
		t.Errorf("New point was incorrect, got: %f, want: %f.", new_points, expected_new_points)
	}
}

func TestRescaleBoundArray_Expand(t *testing.T) {
	old_points  := []float64{0.0,2.5,10.0/3.0,5.0,10.0}
	new_points := RescaleArrayToBounds(0.0, 100.0, old_points)
	expected_new_points  := []float64{0.0,25.0,100.0/3.0, 50.0,100.0}
	
	if (!reflect.DeepEqual(new_points, expected_new_points)) {
		t.Errorf("New point was incorrect, got: %f, want: %f.", new_points, expected_new_points)
	}
}
