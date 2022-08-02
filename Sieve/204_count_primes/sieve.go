package sieve

func countPrimes(n int) int {
	var cnt int
	if n > 2 {
		cnt++
	}
	mark := make([]bool, n, n)
	for i := 3; i*i <= n; i += 2 {
		if !mark[i] {
			for j := i * i; j < n; j += i {
				mark[j] = true
			}
		}
	}
	for i := 3; i < n+1; i += 2 {
		if i < len(mark) && !mark[i] {
			cnt++

		}
	}
	return cnt
}
