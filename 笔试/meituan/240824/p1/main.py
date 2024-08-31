from math import inf

a, b, c, d = map(int, input().split())
n = int(input())
poses = []
for i in range(n):
    poses.append(map(int, input().split()))


total_cost = 0
for p in poses:
    x, y = p
    cur_cos = (abs(x - c) + abs(y - d)) * 2
    total_cost += cur_cos

ans = inf
for p in poses:
    x, y = p
    cur_cos = (abs(x - c) + abs(y - d)) * 2
    remain_cos = total_cost - cur_cos
    new_cos = abs(a - x) + abs(b - y) + abs(x - c) + abs(y - d) + remain_cos
    ans = min(ans, new_cos)

print(ans)