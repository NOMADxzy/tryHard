from logging import lastResort

n, k = map(int, input().split())
MOD = 1000000007

# 6, 2
dp = {}
def dfs(remain: int, down: bool, last: int) -> int:
    if remain == 0:
        return 1
    key = n * remain + last
    if down:
        key *= -1
    if key in dp:
        return dp[key]

    cnt = 0
    if down:
        if last == 1:
            return 0

        l, r = 1, min(remain, last-1)
        for i in range(l, r+1):
            cnt = (cnt + dfs(remain-i, False, i)) % MOD
    else:
        if remain<=last:
            return 0
        l, r = last+1, remain
        for i in range(l, r+1):
            cnt = (cnt + dfs(remain-i, True, i)) % MOD

    dp[key] = cnt
    return cnt

if n-k>0:
    print((dfs(n - k, True, k) + dfs(n - k, False, k)) % MOD)
elif n-k==0:
    print(1)
else:
    print(0)
