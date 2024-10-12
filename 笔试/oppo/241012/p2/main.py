# 0开始a秒弹一次广告，最小相隔b秒关闭一次广告，n时间内最多可以关闭多少广告

t = int(input())

for _ in range(t):
    n, a, b = map(int, input().split())
    ans = 0
    if b <= a:
        ans = n // a
    else:
        ans = min(n // a, n // b)

    print(ans)