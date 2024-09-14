
r, g = map(int, input().split())

max_idx = 0
tmp = []
for _ in range(r):
    x, y, n = map(int, input().split())
    tmp.append((x, y, n))
    max_idx = max(max_idx, x, y)

parents = [[i, 0] for i in range(max_idx+1)]

def find(i: int):
    if parents[i][0] != i:
        f,w = find(parents[i][0])
        parents[i][0], parents[i][1] = f, parents[i][1] + w
    return parents[i]

def union(i1: int, i2: int, w: int):
    f2, w2 = find(i2)
    f1, w1 = find(i1)
    parents[f1] = [f2, w + w2 - w1]

for x, y, n  in tmp:
    f1, w1 = find(x)
    f2, w2 = find(y)
    if f1 != f2:
        union(x, y, n)

# for x, y, n  in tmp:
#     print(find(x)[1] - find(y)[1])

ans = 0
for _ in range(g):
    a, b, m = map(int, input().split())
    f1, w1 = find(a)
    f2, w2 = find(b)

    if f1 == f2 and w1 - w2 == m:
        ans += 1

print(ans)

