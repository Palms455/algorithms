from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        mask = 0
        result_nums: List[List[int]] = []
        while mask < (1 << len(nums)):
            sub_nums = []
            for i in range(len(nums)):
                if (mask & (1 << i)):
                    sub_nums.append(nums[i])
            result_nums.append(sub_nums)
            mask += 1
        return result_nums
