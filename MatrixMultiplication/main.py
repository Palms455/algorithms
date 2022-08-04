from typing import List


def multiplication(A: List[List[int]], B: List[List[int]]) -> List[List[int]]:
    n, m, k = len(A), len(A[0]), len(B[0])
    C = [[0 for i in range(k)] for j in range(n)]
    for i in range(n):
        for j in range(k):
            for l in range(m):
                C[i][j] += A[i][l] * B[l][j]
    return C


if __name__ == "__main__":
    A = [[1, 2, 3], [3, 1, 2]]
    B = [[1, 2], [3, 2], [1, 2]]
    C = [[10, 12], [8, 12]]
    assert multiplication(A, B) == C
