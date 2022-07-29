from typing import List

def sieve(n: int) -> List[int]:
    mark = [True for i in range(n)]
    primes = []
    i = 3
    while i * i <= n:
        if mark[i]:
            j = i * i
            while j < n:
                mark[j] = False
                j += i
        i += 2
    for i in range(3, n+1, 2):
        if mark[i]:
            primes.append(i)
    return primes



if __name__=="__main__":
    print(sieve(100))