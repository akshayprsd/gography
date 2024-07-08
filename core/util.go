package core

import "math"

func FormatNum(num float64, precision int) float64 {
	if precision < 0 {
		return num
	}
	res := math.Pow(10, float64(precision))
	return math.Round(num*res) / res
}

func ToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func ToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}
