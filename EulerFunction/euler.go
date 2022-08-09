package EulerFunction

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// slowEuler time complexity O(n*log(n))
func slowEuler(n int) int {
	count := 1
	for i := 2; i < n; i++ {
		if gcd(i, n) == 1 {
			count++
		}
	}
	return count
}

// fastEuler time complexity O(sqrt(n))
func fastEuler(n int) int {
	result := n
	for i := 2; n >= i*i; i++ {
		if n%i == 0 {
			result = result / i * (i - 1)
		inner:
			for {
				if n%i != 0 {
					break inner
				}
				n /= i
			}
		}
	}
	if n != 1 {
		result = result / n * (n - 1)
	}
	return result
}
