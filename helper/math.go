package helper

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

func Gcd64(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs64(a)
}

func Abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
