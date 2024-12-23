
a = "abcdeg"
b = "axbcdfg"

m = len(a)
n = len(b)

dp1 = [[0]*(n+1) for _ in range(m+1)]
dp2 = [[0]*(n+1) for _ in range(m+1)]
ans1 = 0
ans2 = 0

# 子串
for i in range(m):
    for j in range(n):
        if a[i] == b[j]:
            dp1[i+1][j+1] = 1 + dp1[i][j]
        ans1 = max(ans1, dp1[i+1][j+1])

# 子序列
for i in range(m):
    for j in range(n):
        if a[i] == b[j]:
            dp2[i+1][j+1] = 1 + dp2[i][j]
        else:
            dp2[i+1][j+1] = max(dp2[i][j+1], dp2[i+1][j])
        ans2 = max(ans2, dp2[i+1][j+1])

print(ans1, ans2)
