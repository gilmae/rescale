package rescale

import (
	"math/big"
)

func RescaleArrayToBounds(bound_start float64, bound_end float64, points []float64) []float64 {
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
}

func Rescale(line_start big.Float, line_end big.Float, length int, point int) big.Float {
	step := new (big.Float)
	step.Sub(&line_end, &line_start)
	step.Quo(step,  big.NewFloat(float64(length)))

	half := new(big.Float)
	half.SetFloat64(0.5)

	line_focal_point := new(big.Float)
	line_focal_point.Add(&line_end, &line_start)
	line_focal_point.Mul(line_focal_point, half)

	half_length := new(big.Float)
	half_length.Mul(big.NewFloat(float64(length)), half)
	
	line_focal_point.Add(line_focal_point, big.NewFloat(float64(point)))
	line_focal_point.Sub(line_focal_point, half_length)
	line_focal_point.Mul(line_focal_point, step)
	 
	return *line_focal_point
}
  
func GetZoomedBounds(original_start float64, original_end float64, focal_point float64, zoom float64) (float64,float64) {
	original_width := original_end - original_start
	original_width_to_midpoint := original_width/2.0

	new_width_to_mid_point := original_width_to_midpoint / zoom

	return focal_point-new_width_to_mid_point, focal_point+new_width_to_mid_point
}