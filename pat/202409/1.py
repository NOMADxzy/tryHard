
m, n = map(int, input().split())
ans = 0
cnts = [[0]*n for _ in range(m)]
dirs = [(1, 0), (-1, 0), (0, -1), (0, 1)]

for i in range(m):
    true_i = m-1-i
    row = list(map(int, input().split()))
    for j in range(n):
        dy, dx = dirs[row[j]-1]
        ny, nx = true_i + dy, j + dx
        if nx>=0 and nx<n and ny>=0 and ny<m:
            cnts[ny][nx] += 1

for i in range(m):
    for j in range(n):
        if cnts[i][j] > 1:
            ans += 1

print(ans)