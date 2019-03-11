package even

func Loop(n uint64) (result uint64) {
	result = 1
	var i uint64 = 1
	for ; i <= n; i++ {
		result *= i
	}
	return result
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
