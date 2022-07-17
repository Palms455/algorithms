class Solution:
    def longestValidParentheses(self, s: str) -> int:
        stack = [-1]
        length = 0
        for idx in range(len(s)):
            if s[idx] == "(":
                stack.append(idx)
                continue
            if len(stack) == 1:
                stack[0] = idx
                continue
            stack.pop()
            length = max(length, idx - stack[-1])
        return length


if __name__ == "__main__":
    solution = Solution()
    assert solution.longestValidParentheses("()()") == 4
    assert solution.longestValidParentheses(")()(())") == 6