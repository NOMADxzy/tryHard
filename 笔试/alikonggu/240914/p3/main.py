# n个有权值的线段，找一条垂直线，使他穿过的所有线段的权值 的 或运算结果最大
# 输出这个最大结果

n = int(input())

# lines = []
points_map = {}
points = []
in_w = {}
out_w = {}

def add_point(p):
    if p not in points_map:
        in_w[p] = []
        out_w[p] = []
        points_map[p] = True
        points.append(p)
for _ in range(n):
    l, r, a = map(int, input().split())
    r += 1
    # lines.append((l, r, a))
    add_point(l)
    add_point(r)
    in_w[l].append(a)
    out_w[r].append(a)

points.sort()

# print(in_w)
# print(out_w)
# print(points)

ans = 0
cnts = [0 for _ in range(32)]
for p in points:
    for x in in_w[p]:
        mask = 1
        i = 0
        while mask<=x:
            if x&mask>0 :
                cnts[i] += 1
            i += 1
            mask *= 2
    for x in out_w[p]:
        mask = 1
        i = 0
        while mask<=x:
            if x&mask>0 :
                cnts[i] -= 1
            i += 1
            mask *= 2
    mask = 1
    cur = 0
    for i in range(32):
        if cnts[i] > 0:
            cur += mask
        mask *= 2
    ans = max(ans, cur)

print(ans)
