import random
from typing import List


# Time complexity O(n^2), Memory used O(1)
# At most cases time complexity O(n*log(n))

def partition(nums: List[int], i: int, j: int) -> int:
    #  it might be pivot = i, but for almost sorted array increased time complexity for O(n^2)
    pivot = random.randint(i, j - 1)
    nums[pivot], nums[i] = nums[i], nums[pivot]
    low_index = pivot_point = i
    for k in range(i + 1, j):
        if nums[k] < nums[pivot_point]:
            low_index += 1
            if low_index != k:
                nums[k], nums[low_index] = nums[low_index], nums[k]
    nums[low_index], nums[pivot_point] = nums[pivot_point], nums[low_index]
    return low_index


def quick_sort(nums: List[int], i: int, j: int):
    if i == j:
        return
    pivot_point = partition(nums, i, j)
    quick_sort(nums, i, pivot_point)
    quick_sort(nums, pivot_point + 1, j)


def custom_sort(nums: List[int]):
    quick_sort(nums, 0, len(nums))
    return nums


if __name__ == "__main__":
    lst = [random.randint(1, 1000) for i in range(100)]
    custom_lst = custom_sort(lst.copy())
    assert custom_lst == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_lst}"
    print("sort_succes")
