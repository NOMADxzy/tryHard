from copy import deepcopy
present = list(map(int, input().split(',')))
future = list(map(int, input().split(',')))
budget = int(input())

dp = {0:0}
n = len(present)
ans = 0

for i in range(n):
    newDp = deepcopy(dp)
    cur_cost, cur_value = present[i], future[i]
    for cost, value in dp.items():
        c = cost+cur_cost
        if cost+cur_cost <= budget:
            newDp[c] = max(newDp[c], value+cur_value) if c in newDp else value+cur_value
            ans = max(ans, newDp[c])
    dp = newDp
print(ans)