from random import randint
from typing import List


# Time complexity O(nlog(n))

def merge(left: List[int], right: List[int]) -> List[int]:
    left_idx, right_idx = 0, 0
    new_list = []
    while left_idx < len(left) and right_idx < len(right):
        if left[left_idx] < right[right_idx]:
            new_list.append(left[left_idx])
            left_idx += 1
        else:
            new_list.append(right[right_idx])
            right_idx += 1
    new_list.extend(left[left_idx:])
    new_list.extend(right[right_idx:])
    return new_list


def merge_sort(nums: List[int]) -> List[int]:
    if len(nums) <= 1:
        return nums
    left = merge_sort(nums[:len(nums) // 2])
    right = merge_sort(nums[len(nums) // 2:])
    return merge(left, right)


if __name__ == "__main__":
    lst = [randint(1, 1000) for i in range(100)]
    custom_sort = merge_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
    print("sort_succes")
