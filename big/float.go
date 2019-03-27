package rescale

import (
	"math/big"
)

/*func RescaleArrayToBounds(bound_start float64, bound_end float64, points []float64) []float64 {
	old_min := points[0]
	old_max := points[len(points) - 1]

	original_length := old_max - old_min
	new_length := bound_end - bound_start

	ratio := new_length / original_length 
	
	mapped := make([]float64, len(points))
  	for i, v := range points {
		 mapped[i] = (v - old_min) * ratio + bound_start
  	}
  	return mapped
}*/

func Rescale(line_start big.Float, line_end big.Float, length int, point int) big.Float {
	step := new (big.Float)
	prec := line_start.Prec()

	if prec < line_end.Prec() {
		prec = line_end.Prec()
	}

	step.Sub(&line_end, &line_start)
	step.Quo(step,  big.NewFloat(float64(length)))

	half := big.NewFloat(0.5).SetPrec(prec)

	line_focal_point := new(big.Float).SetPrec(prec)
	line_focal_point.Add(&line_end, &line_start)
	line_focal_point.Mul(line_focal_point, half)

	half_length := new(big.Float).SetPrec(prec)
	half_length.Mul(big.NewFloat(float64(length)), half)
	
	result := new(big.Float).SetPrec(prec)
	result.Sub(big.NewFloat(float64(point)), half_length)
	result.Mul(result, step)
	result.Add(line_focal_point, result)
	 
	return *result
}
  
func GetZoomedBounds(original_start *big.Float, original_end *big.Float, focal_point *big.Float, zoom float64) (big.Float,big.Float) {
	prec := original_start.Prec()

	if prec < original_end.Prec() {
		prec = original_end.Prec()
	}
	
	original_width := new(big.Float).SetPrec(prec)
	original_width.Sub(original_end, original_start)

	half := big.NewFloat(0.5).SetPrec(prec)
	original_width_to_midpoint := new(big.Float).SetPrec(prec)
	original_width_to_midpoint.Mul(original_width, half)

	new_width_to_mid_point := new(big.Float).SetPrec(prec)
	new_width_to_mid_point.Quo(original_width_to_midpoint, big.NewFloat(zoom).SetPrec(prec))

	new_start := new(big.Float).SetPrec(prec)
	new_end := new(big.Float).SetPrec(prec)
	return *new_start.Sub(focal_point, new_width_to_mid_point), *new_end.Add(focal_point, new_width_to_mid_point)
}