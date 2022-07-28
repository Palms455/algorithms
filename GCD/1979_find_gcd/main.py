from typing import List

class Solution:
    def findGCD(self, nums: List[int]) -> int:
        return self.gcd(max(nums), min(nums))

    def gcd(self, a: int, b: int) -> int:
        return a if b == 0 else self.gcd(b, a % b)
