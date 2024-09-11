
n, m = map(int, input().split())
MOD = 1000000007

dp = {}


def dfs(remain0, remain1, pos) -> int:
    key = 1000020001 * pos + 10001 * remain0 + remain1
    if key in dp:
        return dp[key]
    limit = min(remain0, remain1)
    cnt = 1
    max_0 = remain0 // pos
    max_1 = remain1 // pos
    for i in range(1, max_0 + 1):
        for j in range(1, max_1 + 1):
            # print(i, j)
            cnt = (cnt + dfs(remain0 - i * pos, remain1 - j * pos, pos + 1)) % MOD

    dp[key] = cnt
    return cnt


print(dfs(n, m, 1))
