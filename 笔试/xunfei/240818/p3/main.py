
n = int(input())
arr = list(map(int, input().split()))
# 3
# 2 1 4

pre_sums = [0 for _ in range(n+1)]
pre_sums[0] = arr[0]
for i in range(1, n+1):
    pre_sums[i] = pre_sums[i-1] + arr[i-1]

dp = {}
def dfs(l: int, r: int) -> int:
    key = n*l + r
    if key in dp:
        return dp[key]

    cur_sum = pre_sums[r+1] - pre_sums[l]
    ret = cur_sum
    if l-1>=0 and arr[l-1]>cur_sum:
        ret = max(ret, dfs(l-1, r))
    if r+1<n and arr[r+1]>cur_sum:
        ret = max(ret, dfs(l, r+1))

    dp[key] = ret
    return ret

ans = []
for i in range(n):
    ans.append(dfs(i, i))

print(" ".join(map(str, ans)))
