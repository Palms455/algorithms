package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == -1 {
		return myPow(x, n+1) / x
	}
	if n%2 == 1 {
		return myPow(x, n-1) * x
	}
	b := myPow(x, n/2)
	return b * b
}