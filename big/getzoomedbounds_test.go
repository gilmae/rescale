package rescale

import (
	"testing"
	"math/big"
)

func TestGetZoomedBounds(t *testing.T) {

	start := big.NewFloat(0.0).SetPrec(50)
	end := big.NewFloat(10.0).SetPrec(50)
	point := big.NewFloat(5.0).SetPrec(50)

	expected_new_start := big.NewFloat(2.5).SetPrec(50)
	expected_new_end := big.NewFloat(7.5).SetPrec(50)

	new_start, new_end := GetZoomedBounds(start,end,point,2);
	if new_start.Cmp(expected_new_start) != 0 {
		t.Errorf("New minimum bound were incorrect, got: %s, want: %s.", new_start.Text('f', 10), expected_new_start.Text('f', 10))
	}
	if new_end.Cmp(expected_new_end) != 0 {
		t.Errorf("New maximum bound were incorrect, got: %s, want: %s.", new_end.Text('f', 10), expected_new_end.Text('f', 10))
	}
}

func TestGetZoomedBounds_With_UnEvenLength(t *testing.T) {
	start := big.NewFloat(1.0).SetPrec(50)
	end := big.NewFloat(10.0).SetPrec(50)
	point := big.NewFloat(5.0).SetPrec(50)

	expected_new_start := big.NewFloat(2.75).SetPrec(50)
	expected_new_end := big.NewFloat(7.25).SetPrec(50)

	new_start, new_end := GetZoomedBounds(start,end,point,2);
	
	if new_start.Cmp(expected_new_start) != 0 {
		t.Errorf("New minimum bound were incorrect, got: %s, want: %s.", new_start.Text('f', 10), expected_new_start.Text('f', 10))
	}
	if new_end.Cmp(expected_new_end) != 0 {
		t.Errorf("New maximum bound were incorrect, got: %s, want: %s.", new_end.Text('f', 10), expected_new_end.Text('f', 10))
	}
}

func TestGetZoomedBounds_With_GreaterZoom(t *testing.T) {

	start := big.NewFloat(0.0).SetPrec(50)
	end := big.NewFloat(10.0).SetPrec(50)
	point := big.NewFloat(5.0).SetPrec(50)

	expected_new_start := big.NewFloat(4.5).SetPrec(50)
	expected_new_end := big.NewFloat(5.5).SetPrec(50)

	new_start, new_end := GetZoomedBounds(start,end,point,10);
	
	if new_start.Cmp(expected_new_start) != 0 {
		t.Errorf("New minimum bound were incorrect, got: %s, want: %s.", new_start.Text('f', 10), expected_new_start.Text('f', 10))
	}
	if new_end.Cmp(expected_new_end) != 0 {
		t.Errorf("New maximum bound were incorrect, got: %s, want: %s.", new_end.Text('f', 10), expected_new_end.Text('f', 10))
	}
}

func TestGetZoomedBounds_With_NonCentre(t *testing.T) {
	start := big.NewFloat(0.0).SetPrec(50)
	end := big.NewFloat(10.0).SetPrec(50)
	point := big.NewFloat(3.0).SetPrec(50)

	expected_new_start := big.NewFloat(0.5).SetPrec(50)
	expected_new_end := big.NewFloat(5.5).SetPrec(50)

	new_start, new_end := GetZoomedBounds(start,end,point,2);
	
	if new_start.Cmp(expected_new_start) != 0 {
		t.Errorf("New minimum bound were incorrect, got: %s, want: %s.", new_start.Text('f', 10), expected_new_start.Text('f', 10))
	}
	if new_end.Cmp(expected_new_end) != 0 {
		t.Errorf("New maximum bound were incorrect, got: %s, want: %s.", new_end.Text('f', 10), expected_new_end.Text('f', 10))
	}
}