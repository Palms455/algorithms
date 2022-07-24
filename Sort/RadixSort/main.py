from random import randint
from typing import List


# Time complexity O(d * n), Memory used O(n)

def radix_sort(nums: List[int]) -> List[int]:
    power_of_ten = 1
    arr = [[] for i in range(10)]
    for pow in range(5):
        for num in nums:
            arr[int(num/power_of_ten) % 10].append(num)
        nums = []
        for idx, item in enumerate(arr):
            nums.extend(item)
            arr[idx] = []
        power_of_ten *= 10
    return nums


if __name__ == "__main__":
    lst = [randint(1, 10) for i in range(100)]
    custom_sort = radix_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
    print("sort_succes")

