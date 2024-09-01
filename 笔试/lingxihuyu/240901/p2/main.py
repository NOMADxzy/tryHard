
n = int(input())
arr = list(map(int, input().split()))

dp = {}
ans = 0

# 45 55 56

for i in range(n):
    max_len = [0, 1, 0]
    if arr[i]-1 in dp:
        best = max(dp[arr[i]-1][1], dp[arr[i]-1][2])
        max_len[0] = max(max_len[0], best + 1)
    if arr[i] in dp:
        max_len[0] = max(max_len[0], dp[arr[i]][0] + 1) if dp[arr[i]][0]>0 else 0
        max_len[1] = max(max_len[1], dp[arr[i]][1] + 1)
        max_len[2] = max(max_len[2], dp[arr[i]][2] + 1) if dp[arr[i]][2]>0 else 0
    if arr[i] + 1 in dp:
        best = max(dp[arr[i]+1][0], dp[arr[i]+1][1])
        max_len[2] = max(max_len[2], best + 1)
    dp[arr[i]] = max_len
    ans = max(ans, max_len[0], max_len[1], max_len[2])

print(ans)
