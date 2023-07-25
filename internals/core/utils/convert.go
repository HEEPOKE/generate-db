package utils

import "strconv"

func ConvertToFloat32(s string) float32 {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	f32 := float32(f64)

	return f32
}
