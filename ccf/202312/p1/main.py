

n, m = map(int, input().split())
# 4 2
# 0 0
# -1 -1
# 1 2
# 0 -1

features = []

for i in range(n):
    feature = list(map(int, input().split()))
    features.append(feature)


for i in range(n):
    best = 0
    for j in range(n):
        if i == j:
            continue
        ok = True
        c, f = features[i], features[j]
        for k in range(m):
            if c[k] >= f[k]:
                ok = False
                break
        if ok:
            best = j + 1
            break
    print(best)




