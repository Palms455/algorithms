import collections


class Solution:
    def countGoodSubstrings(self, s: str) -> int:
        left, nums, state = 0, 0, collections.defaultdict(int)

        for idx, elem in enumerate(s):
            state[elem] += 1

            if idx - left == 3:
                state[s[left]] -= 1
                if state[s[left]] == 0:
                    del (state[s[left]])
                left += 1

            if len(state) == 3:
                nums += 1
        return nums
