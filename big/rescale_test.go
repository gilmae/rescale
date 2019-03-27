package rescale

import (
	"testing"
	"math/big"
)

func TestRescale_Shrink(t *testing.T) {
	new_point := Rescale(*big.NewFloat(0.0), *big.NewFloat(10.0), 100, 50)
	expected_new_point := big.NewFloat(5.0)
	
	if (new_point.Cmp(expected_new_point) != 0) {
		t.Errorf("New point was incorrect, got: %s, want: %s.", new_point.Text('e', 20), expected_new_point.Text('e', 20))
	}
}

func TestRescale_Expand(t *testing.T) {
	new_point := Rescale(*big.NewFloat(0.0), *big.NewFloat(100.0), 10, 5)
	expected_new_point := big.NewFloat(50.0)
	
	if (new_point.Cmp(expected_new_point) != 0) {
		t.Errorf("New point was incorrect, got: %s, want: %s.", new_point.Text('e', 20), expected_new_point.Text('e', 20))
	}
}
