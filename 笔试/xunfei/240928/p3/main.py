
n, m, q = map(int, input().split())
grid = [[0] * n for _ in range(n)]
for _ in range(m):
    u, v = map(int, input().split())
    u -= 1
    v -= 1
    grid[u][v] += 1

for k in range(n):
    for i in range(n):
        for j in range(n):
            new_dist = min(grid[i][k], grid[k][j])
            grid[i][j] = max(grid[i][j], new_dist)

for _ in range(q):
    u, v = map(int, input().split())
    u -= 1
    v -= 1
    print(grid[u][v])
