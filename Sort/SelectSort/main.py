from random import randint
from typing import List


# Time complexity O(n), Memory used O(1)

def select_sort(nums: List[int]) -> List[int]:
    for i in range(len(nums)):
        min_idx = i
        for j in range(i + 1, len(nums)):
            if nums[min_idx] > nums[j]:
                nums[min_idx], nums[j] = nums[j], nums[min_idx]
    return nums


if __name__ == "__main__":
    lst = [randint(1, 100) for i in range(100)]
    custom_sort = select_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
