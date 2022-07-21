package main

func numberOfSteps(num int) int {
	steps := 0
	for ; num > 0; steps++ {
		if (num & 1) == 1 {
			num--
			continue
		}
		num = num >> 1
	}
	return steps
}
