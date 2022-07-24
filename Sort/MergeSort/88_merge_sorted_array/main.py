from typing import List

class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums = nums1[:m]
        nums1 = []
        left = right = 0
        while left < len(nums) and right < len(nums2):
            if nums[left] < nums2[right]:
                nums1.append(nums[left])
            else:
                nums1.append(nums2[right])
        nums1.extend(nums[left:])
        nums1.extend(nums[right:])