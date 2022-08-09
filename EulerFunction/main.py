def gcd(a: int, b: int) -> int:
    return a if b == 0 else gcd(b, a % b)


def slow_euler(n: int) -> int:
    count = 1
    for i in range(2, n):
        if gcd(i, n) == 1:
            count += 1
    return count


def fast_euler(n: int) -> int:
    result = n
    prime = 2
    while n >= prime * prime:
        if n % prime == 0:
            result = result / prime * (prime - 1)
            while n % prime == 0:
                n /= prime
        prime += 1
    if n != 1:
        result = result / n * (n - 1)
    return result

if __name__ == "__main__":
    assert slow_euler(9) == 6
    assert fast_euler(9) == 6
    assert slow_euler(36) == 12
    assert fast_euler(36) == 12
    assert slow_euler(84) == 24
    assert fast_euler(84) == 24
