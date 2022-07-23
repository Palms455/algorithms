from random import randint
from typing import List


# Time complexity O(n^2), Memory used O(1)


def insert_sort(nums: List[int]) -> List[int]:
    for i in range(1, len(nums)):
        cur = nums[i]
        j = i - 1
        while j >= 0 and nums[j] > cur:
            nums[j+1] = nums[j]
            j -= 1
        nums[j+1] = cur
    return nums


if __name__ == "__main__":
    lst = [randint(1, 1000) for i in range(100)]
    custom_sort = insert_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
    print("sort_succes")

