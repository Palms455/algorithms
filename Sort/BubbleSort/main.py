from random import randint
from typing import List


# Time complexity O(n^2), Memory used O(1)

def bubble_sort(nums: List[int]) -> List[int]:
    for i in range(len(nums)):
        is_sorted = True
        for j in range(1, len(nums) - i):
            if nums[j] < nums[j - 1]:
                is_sorted = False
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
        if is_sorted:
            return nums
    return nums


if __name__ == "__main__":
    lst = [randint(1, 1000) for i in range(100)]
    custom_sort = bubble_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
    print("sort_succes")

