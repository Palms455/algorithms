from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            m = (left + right) // 2
            if nums[m] == target:
                return m
            if nums[m] > target:
                right = m - 1
                continue
            left = m + 1
        return -1


if __name__ == "__main__":
    solution = Solution()
    lst = [1, 3, 4, 8, 9, 12]
    assert solution.search(lst, 8) == 3
    assert solution.search(lst, 3) == 1
    assert solution.search(lst, 2) == -1
