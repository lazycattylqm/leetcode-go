package leetcode

func convertTemperature(celsius float64) []float64 {

	v1 := celsius + 273.15
	v2 := celsius*1.80 + 32.00
	return []float64{v1, v2}
}
