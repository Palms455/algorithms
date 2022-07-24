from typing import List


class Solution:
    def maximumGap(self, nums: List[int]) -> int:
        power_of_ten = 1
        for _ in range(9):
            arr = [[] for i in range(10)]
            for num in nums:
                arr[int(num / power_of_ten) % 10].append(num)
            nums = []
            for elem in arr:
                nums.extend(elem)
            power_of_ten *= 10
        cnt = 0
        for i in range(1, len(nums)):
            cnt = max(cnt, nums[i] - nums[i - 1])
        return cnt
