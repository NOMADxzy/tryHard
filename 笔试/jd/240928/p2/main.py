
m, n = map(int, input().split())
# 2 3
# 3 3 2
# 3 2 1
heights = []

for _ in range(m):
    heights.append(list(map(int, input().split())))

col_max, row_max = [0]*n, [0]*m
for i in range(m):
    for j in range(n):
        row_max[i] = max(row_max[i], heights[i][j])
        col_max[j] = max(col_max[j], heights[i][j])

v1, v2, v3 = sum(col_max), sum(row_max), m*n
print(v1, v2, v3)
