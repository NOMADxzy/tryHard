#

n = int(input())
arrs = []
for _ in range(3):
    arrs.append(list(map(int, input().split())))


dp = [[0]*3 for _ in range(n)]

for i in range(1, n):
    for j in range(3):
        a = dp[i-1][0] + abs(arrs[0][i-1] - arrs[j][i])
        b = dp[i-1][1] + abs(arrs[1][i-1] - arrs[j][i])
        c = dp[i-1][2] + abs(arrs[2][i-1] - arrs[j][i])
        dp[i][j] = min(a, b, c)

# print(dp)
print(min(dp[-1][0], dp[-1][1], dp[-1][2]))
