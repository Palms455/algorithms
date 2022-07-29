class Solution:
    def countPrimes(self, n: int) -> int:
        cnt = 1 if n > 2 else 0
        mark = [True for i in range(n)]
        i = 3
        while i * i <= n:
            if mark[i]:
                j = i * i
                while j < n:
                    mark[j] = False
                    j += i
            i += 2
        for i in range(3, n + 1, 2):
            if i < len(mark) and mark[i]:
                cnt += 1
        return cnt
