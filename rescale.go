package rescale

func RescaleArrayToBounds(bound_start float64, bound_end float64, points []float64) []float64 {
	old_min := points[0]
	old_max := points[len(points) - 1]

	original_length := old_max - old_min
	new_length := bound_end - bound_start

	ratio := original_length / new_length
	
	mapped := make([]float64, len(points))
  for i, v := range points {
		 mapped[i] = bound_start + ((v - old_min) / ratio)
  }
  return mapped
}

func Rescale(line_start float64, line_end float64, length int, point int) float64 {
	step := (line_end - line_start) / float64(length)
	line_focal_point := (line_end + line_start) / 2.0
	 
	return line_focal_point + (float64(point) - (float64(length)/2.0)) * step
}
  
func GetZoomedBounds(original_start float64, original_end float64, focal_point float64, zoom float64) (float64,float64) {
	original_width := original_end - original_start
	original_width_to_midpoint := original_width/2.0

	new_width_to_mid_point := original_width_to_midpoint / zoom

	return focal_point-new_width_to_mid_point, focal_point+new_width_to_mid_point
}