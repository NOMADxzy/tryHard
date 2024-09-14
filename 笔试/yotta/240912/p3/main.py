import sys

n = int(input())
lines = [x[:-1] for x in sys.stdin.readlines()]

# 一辆单向货车，有n个任务，每个任务从f到t可以获得t-f的基础奖励+r的额外奖励，可以在同一个地点卸货装货，求能获得的最大奖励

pres = [[] for _ in range(n)]
for line in lines:
    f, t, r = map(int, input().split())
    f -= 1
    t -= 1
    reward = t - f + r
    pres[t].append((f, reward))

dp = [0 for _ in range(n)]

for i in range(n):
    dp[i] = dp[i-1] if i-1>=0 else 0
    for pre_pos, reward in pres[i]:
        dp[i] = max(dp[i], dp[pre_pos] + reward)

print(dp[-1])
