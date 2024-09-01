import sys


n, m = map(int, input().split())
costs = list(map(int, sys.stdin.readline().split()))
energys = list(map(int, sys.stdin.readline().split()))

# 2 3
# 10 15
# 10 15 20

costs.sort()
energys.sort()

ans = 0
i,j = 0,0

while i<n and j<m:
    while j<m and energys[j] < costs[i]:
        j += 1
    if j==m:
        break

    ans += 1
    i += 1
    j += 1

print(ans)