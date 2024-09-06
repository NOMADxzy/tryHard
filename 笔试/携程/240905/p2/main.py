
n = int(input())
s = input()

#00001

ans = 0
dp = [[0,0] for _ in range(n)]
for i in range(n):
    if s[i] == "1":
        dp[i][0] = 1
    else:
        dp[i][1] = 1
        ans += 1

for delta in range(1, n):
    newDp = [[0,0] for _ in range(n-delta)]
    for i in range(0, n-delta):
        l, r = i, i+delta
        if s[r] == '0':
            newDp[l][0] = min(dp[l][0], dp[l][1]+1)
            newDp[l][1] = min(dp[l][0]+1, dp[l][1]+2)
        else:
            newDp[l][1] = min(dp[l][1], dp[l][0]+1)
            newDp[l][0] = min(dp[l][1]+1, dp[l][0]+2)
        if  (r-l+1) % 2 == 1 and newDp[l][1] % 2 == 1:
            # print(s[l:r+1])
            ans += 1
    dp = newDp

print(ans)
