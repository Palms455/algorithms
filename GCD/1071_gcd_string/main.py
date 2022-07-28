class Solution:
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        gcd_val = self.gcd(len(str1), len(str2))
        lcm_val = self.lcm(len(str1), len(str2), gcd_val)
        for i in range(lcm_val):
            if str1[i % len(str1)] != str2[i % len(str1)]:
                return ""
        return str1[:gcd_val]



    def gcd(self, a: int, b: int) -> int:
        if b == 0:
            return a
        return self.gcd(b, a%b)

    def lcm(self, a, b, gcd) -> int:
        return a / gcd * b