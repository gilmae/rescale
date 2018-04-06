package rescale

import "testing"

func TestGetZoomedBounds(t *testing.T) {
	new_start, new_end := GetZoomedBounds(0.0,10.0,5,2);
	if new_start != 2.5 {
		t.Errorf("New minimum bound were incorrect, got: %f, want: %f.", new_start, 2.5)
	}
	if new_end != 7.5 {
		t.Errorf("New maximum bound were incorrect, got: %f, want: %f.", new_end, 7.5)
	}
}

func TestGetZoomedBounds_With_UnEvenLength(t *testing.T) {
	new_start, new_end := GetZoomedBounds(1.0,10.0,5,2);
	expected_new_start := 2.75
	expected_new_end := 7.25
	
	if new_start != expected_new_start {
		t.Errorf("New minimum bound were incorrect, got: %f, want: %f.", new_start, expected_new_start)
	}
	if new_end != expected_new_end {
		t.Errorf("New maximum bound were incorrect, got: %f, want: %f.", new_end, expected_new_end)
	}
}

func TestGetZoomedBounds_With_GreaterZoom(t *testing.T) {
	new_start, new_end := GetZoomedBounds(0.0,10.0,5,10);
	expected_new_start := 4.5
	expected_new_end := 5.5
	
	if new_start != expected_new_start {
		t.Errorf("New minimum bound were incorrect, got: %f, want: %f.", new_start, expected_new_start)
	}
	if new_end != expected_new_end {
		t.Errorf("New maximum bound were incorrect, got: %f, want: %f.", new_end, expected_new_end)
	}
}

func TestGetZoomedBounds_With_NonCentre(t *testing.T) {
	new_start, new_end := GetZoomedBounds(0.0,10.0,3,2);
	expected_new_start := 0.5
	expected_new_end := 5.5
	
	if new_start != expected_new_start {
		t.Errorf("New minimum bound were incorrect, got: %f, want: %f.", new_start, expected_new_start)
	}
	if new_end != expected_new_end {
		t.Errorf("New maximum bound were incorrect, got: %f, want: %f.", new_end, expected_new_end)
	}
}