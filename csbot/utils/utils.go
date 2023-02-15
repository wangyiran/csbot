package utils

func MaxPrice(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MinPrice(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
