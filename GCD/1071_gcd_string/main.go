package main


func gcdOfStrings(str1 string, str2 string) string {
	g := gcd(len(str1), len(str2))
	l := lcm(len(str1), len(str2), g)
	for i:=0; i< l; i++ {
		if str1[i%len(str1)] != str2[i%len(str2)] {
			return ""
		}
	}
	return str1[:g]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b, gcd int) int {
	return a / gcd * b
}