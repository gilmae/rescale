package rescale

func Rescale_Array(line_start float64, line_end float64, line_focal_point float64, total_points int, points []int, zoom float64) []float64 {
  step := (line_end - line_start) / float64(total_points) / zoom
  mapped := make([]float64, len(points))
    for i, v := range points {
      mapped[i] = line_focal_point + (float64(v) - (float64(total_points)/2.0)) * step
    }
    return mapped
}

func Rescale(line_start float64, line_end float64, total_points int, point int) float64 {
	step := (line_end - line_start) / float64(total_points)
	line_focal_point := (line_start + line_end) / 2.0
	 
	return line_focal_point + (float64(point) - (float64(total_points)/2.0)) * step
  }
  
func Get_Zoomed_Bounds(original_start float64, original_end float64, focal_point float64, zoom float64) (float64,float64) {
	step_precision := 1000000.0
	
	original_step := (original_end - original_start) / step_precision
	zoomed_step := original_step / zoom
  
	  
	  new_start := focal_point - zoomed_step * step_precision/2.0
	  new_end := focal_point + zoomed_step * step_precision/2.0
  
	  return new_start, new_end
  }