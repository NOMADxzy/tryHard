import sys

MOD = 1000000007
n, m = map(int, sys.stdin.readline().split())
m = m % MOD

# n, m = map(int, input().split())

ans = 1


left = 1
right = 1
for k in range(1, n+1):
    right = (right * m) % MOD
    left = left * ((n - k + 1) % MOD)
    left //= k
    inc = (left * right) % MOD
    ans = (ans + inc) % MOD

print(ans)
# print(math.pow(2, 29))