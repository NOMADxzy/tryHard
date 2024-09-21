
# 每一行、每一列、每条对角线上的1最多为多少？

# 5 5
# 0 0 0 0 0
# 1 1 0 0 0
# 0 1 1 0 0
# 0 1 0 1 0
# 1 0 0 0 1

m, n = map(int, input().split())
grid = []
for _ in range(n):
    grid.append(list(map(int, input().split())))

ans = 0
for i in range(m):
    x, y = i, 0
    cur = 0
    while 0 <= x < m and 0 <= y < n:
        cur = 0 if grid[x][y] == 0 else cur + 1
        ans = max(cur, ans)
        y += 1

for j in range(n):
    x, y = 0, j
    cur = 0
    while 0 <= x < m and 0 <= y < n:
        cur = 0 if grid[x][y] == 0 else cur + 1
        ans = max(cur, ans)
        x += 1

for i in range(0, m+n-1):
    if i<m:
        x, y = i,0
    else:
        x, y = m-1, i-m+1
    cur = 0
    while 0 <= x < m and 0 <= y < n:
        cur = 0 if grid[x][y] == 0 else cur + 1
        ans = max(cur, ans)
        y += 1
        x -= 1

for i in range(0, m+n-1):
    if i<m:
        x, y = m-1-i,0
    else:
        x, y = 0, i-m+1
    cur = 0
    while 0 <= x < m and 0 <= y < n:
        cur = 0 if grid[x][y] == 0 else cur + 1
        ans = max(cur, ans)
        y += 1
        x += 1

print(ans)