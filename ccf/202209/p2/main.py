from math import inf

n, x = map(int, input().split())

prices = []
for _ in range(n):
    prices.append(int(input()))

# 20 90 60 60

dp = {0}
cur_best = inf

for i in range(0, n):
    newDp = set(dp)
    for e in dp:
        if e + prices[i] < cur_best:
            newDp.add(e + prices[i])
            if e + prices[i] >= x:
                cur_best = min(cur_best, e + prices[i])
    dp = newDp

print(cur_best)

