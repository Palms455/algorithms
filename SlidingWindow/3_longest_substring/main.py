class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        length, left, state = 0, 0, set()
        for idx, elem in enumerate(s):
            while elem in state:
                state.remove(s[left])
                left += 1
            state.add(elem)
            length = max(length, idx - left +1)

        return length
