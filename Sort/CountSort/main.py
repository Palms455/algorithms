from random import randint
from typing import List


# Time complexity O(n+m), Memory used O(1)

def count_sort(nums: List[int]) -> List[int]:
    minimum = nums[0]
    maximum = nums[-1]

    for num in nums:
        if num < minimum:
            minimum = num
        if num > maximum:
            maximum = num
    cnt = [0 for i in range(minimum, maximum+1)]
    for num in nums:
        cnt[num-minimum] += 1
    new = []
    for idx, num in enumerate(cnt):
        for i in range(num):
            new.append(idx+minimum)
    return new


if __name__ == "__main__":
    lst = [randint(1, 1000) for i in range(100)]
    custom_sort = count_sort(lst.copy())
    assert custom_sort == sorted(lst), f"Sort error. Expected {sorted(lst)}, given: {custom_sort}"
    print("sort_succes")
    print(lst)
    print(custom_sort)

