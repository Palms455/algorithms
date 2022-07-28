package main

func findGCD(nums []int) int {
	maximum, minimum := nums[0], nums[0]
	for _, num := range nums {
		if maximum < num {
			maximum = num
		}
		if minimum > num {
			minimum = num
		}
	}
	return recursiveGCD(maximum, minimum)
}

func recursiveGCD(a, b int) int {

	if b == 0 {
		return a
	}
	return recursiveGCD(b, a%b)
}

func iterativeGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
