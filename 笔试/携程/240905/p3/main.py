from os import stat

n, m, k = map(int, input().split())

k_arr = []
tmp_val = k
while tmp_val > 0:
    k_arr.append(tmp_val % 10)
    tmp_val //= 10
while len(k_arr) < m:
    k_arr.append(0)
k_arr.reverse()
# print(k_arr)
# used = [False for _ in range(n+1)]
state = 0
dp = {}


def dfs(pos: int, limit: bool, state: int) -> int:
    if pos == m:
        return 0 if limit else 1
    key = state
    if limit:
        key = -key
    if key in dp:
        return dp[key]
    cnt = 0

    for v in range(0, n + 1):
        if (limit and v < k_arr[pos]) or state & (1 << v) > 0:
            continue
        if pos == 0 and v == 0 and m>1:
            continue
        cnt += dfs(pos + 1, limit and v == k_arr[pos], state + 1 << v)

    dp[key] = cnt
    return cnt


print(dfs(0, True, 0))
# state + (1 << v)