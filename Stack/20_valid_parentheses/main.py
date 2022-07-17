class Solution:
    pairs = {
        "{": "}",
        "(": ")",
        "[": "]"
    }
    def is_open(self, symbol: str) -> bool:
        return symbol in ("({[")

    def is_pair(self, open_s: str, close_s: str) -> bool:
        return self.pairs.get(open_s) == close_s

    def isValid(self, s: str) -> bool:
        open_list = []
        for symbol in s:
            if self.is_open(symbol):
                open_list.append(symbol)
                continue
            if not open_list or not self.is_pair(open_list.pop(), symbol):
                return False
        return not bool(open_list)


if __name__ == "__main__":
    solution = Solution()
    assert solution.isValid("()") is True
    assert solution.isValid("()[]{}") is True
    assert solution.isValid("(]") is False
    print("all check passed")