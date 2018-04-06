package rescale

import "testing"

func TestRescale_Shrink(t *testing.T) {
	new_point := Rescale(0.0, 10.0, 100, 50)
	expected_new_point := 5.0
	
	if (new_point != expected_new_point) {
		t.Errorf("New point was incorrect, got: %f, want: %f.", new_point, expected_new_point)
	}
}

func TestRescale_Expand(t *testing.T) {
	new_point := Rescale(0.0, 100.0, 10, 5)
	expected_new_point := 50.0
	
	if (new_point != expected_new_point) {
		t.Errorf("New point was incorrect, got: %f, want: %f.", new_point, expected_new_point)
	}
}
