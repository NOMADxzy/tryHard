
n = int(input())
s = input()

# 4
# RRGR

dp = [1-int(s[0] == 'R'),1-int(s[0] == 'G'),1-int(s[0] == "B")]

for i in range(1, n):
    newDp = [0,0,0]
    if s[i] == "R":
        newDp[0] = min(dp[1], dp[2])
        newDp[1] = dp[1] + 1
        newDp[2] = dp[2] + 1
    elif s[i] == "G":
        newDp[1] = min(dp[0], dp[2])
        newDp[0] = dp[0] + 1
        newDp[2] = dp[2] = 1
    if s[i] == "B":
        newDp[2] = min(dp[0], dp[1])
        newDp[1] = dp[1] + 1
        newDp[2] = dp[2] = 1

    dp = newDp

print(min(dp))
