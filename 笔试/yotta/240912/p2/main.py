
grid = []
for _ in range(4):
    grid.append(list(map(int, input().split())))

dirs = [[-1,0],[1,0],[0,-1],[0,1], [-1,-1],[-1,1],[1,-1],[1,1]]

visited = [[False for _ in range(4)] for _ in range(4)]

def dfs(x, y, t):
    cnt = 1
    visited[x][y] = True
    for dir in dirs:
        nx, ny = x + dir[0], y + dir[1]
        if nx>=0 and nx<4 and ny>=0 and ny<4 and grid[nx][ny] == t and not visited[nx][ny]:
            cnt += dfs(nx, ny, t)
    return cnt

values = [1,3,5]
ans = 0
for i in range(4):
    for j in range(4):
        if not visited[i][j]:
            ans = max(ans, values[grid[i][j]-1] * dfs(i, j, grid[i][j]))

print(ans)