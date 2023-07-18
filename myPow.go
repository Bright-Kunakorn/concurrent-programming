func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}

	return result
}
