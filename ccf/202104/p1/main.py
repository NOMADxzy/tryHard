
n, m, L = map(int, input().split())
grid = []

for _ in range(n):
    grid.append(list(map(int, input().split())))

ans = [0 for _ in range(L)]

for i in range(n):
    for j in range(m):
        ans[grid[i][j]] += 1

print(" ".join(map(str, ans)))