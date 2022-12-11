class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        def partition(nums, left, right):
            pivot = random.randint(left, right - 1)
            nums[pivot], nums[left] = nums[left], nums[pivot]
            low_idx = pivot_point = left
            for i in range(left + 1, right):
                if nums[i] < nums[pivot_point]:
                    low_idx += 1
                    if low_idx != i:
                        nums[i], nums[low_idx] = nums[low_idx], nums[i]
            nums[low_idx], nums[pivot_point] = nums[pivot_point], nums[low_idx]
            return low_idx

        def quicksort(nums, left, right):
            if left == right:
                return
            pivot = partition(nums, left, right)
            quicksort(nums, left, pivot)
            quicksort(nums, pivot + 1, right)

        quicksort(nums, 0, len(nums))
        return nums[-k]
