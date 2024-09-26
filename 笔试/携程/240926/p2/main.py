# 问题：a、b、c、d 四个数的非1最小公约数是多少

T = int(input())


def gcd(a: int, b: int) -> int:  # 获取最大公共因数
    if a < b:
        a, b = b, a
    while b != 0:
        a, b = b, a % b
    return a


def simpler(x: int):  # 获取最小的因子
    for i in range(2, x):
        if i * i > x:
            break
        if x % i == 0:
            return i
    return x


for _ in range(T):
    a, b, c, d = map(int, input().split())

    ans = a
    arr = [b, c, d]
    idx = 0
    while ans > 1 and idx < len(arr):
        ans = gcd(ans, arr[idx])
        idx += 1
    print(simpler(ans) if ans > 1 else -1)
