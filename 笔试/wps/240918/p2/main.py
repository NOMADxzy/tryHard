
n, m = map(int, input().split())

limits = []
for _ in range(n):
    l, r = map(int, input().split())
    limits.append((l,r))

min_right = [0 for _ in range(n)]
min_right[-1] = limits[-1][0]
for i in range(n-2, -1, -1):
    min_right[i] = min_right[i+1] + limits[i][0]
max_right = [0 for _ in range(n)]
max_right[-1] = limits[-1][1]
for i in range(n-2, -1, -1):
    max_right[i] = max_right[i+1] + limits[i][1]

# print(min_right)
# print(max_right)
ans = 0
if min_right[0] <= m:
    dp = {}
    def dfs(remain: int, pos: int) -> int:
        key = remain * n + pos
        if key in dp:
            return dp[key]
        if pos == n-1:
            return 1
        l, r = limits[pos]
        cnt = 0
        l = max(l, remain - max_right[pos+1])
        r = min(r, remain - min_right[pos+1], remain)
        for cur_n in range(l, r+1):
            cnt += dfs(remain - cur_n, pos+1)

        dp[key] = cnt
        return cnt

    ans = dfs(m, 0)

print(ans)