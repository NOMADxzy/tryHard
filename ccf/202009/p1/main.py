n, x0, y0 = map(int, input().split())

min_vals = []


def get_dist(x, y):
    return (x - x0) * (x - x0) + (y - y0) * (y - y0)


for idx in range(n):
    x, y = map(int, input().split())
    dist = get_dist(x, y)
    if len(min_vals) == 0:
        min_vals.append((dist, idx + 1))
    else:
        i = 0
        while i < len(min_vals):
            e = min_vals[i]
            if dist < e[0]:
                min_vals.insert(i, (dist, idx + 1))
                break
            i += 1
        if i==len(min_vals):
            min_vals.append((dist, idx + 1))
        min_vals = min_vals[:3]

for i in range(len(min_vals)):
    print(min_vals[i][1])
