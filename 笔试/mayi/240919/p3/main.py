
# 4x4 矩阵中，已经走了一些步骤，问反悔多少次可以把矩阵走完
# 输出最小的反悔次数

grid = []
for _ in range(4):
    grid.append(list(map(int, input().split())))

dirs = [[-1, 0], [1, 0], [0, -1], [0, 1]]


def dfs(x: int, y: int, step: int) -> bool:
    grid[x][y] = step
    if step == 16:
        return True
    ok = False
    for dir in dirs:
        nx, ny = x + dir[0], y + dir[1]
        if nx >= 0 and nx < 4 and ny >= 0 and ny < 4 and grid[nx][ny] == 0 and dfs(nx, ny, step + 1):
            return True

    grid[x][y] = 0
    return False


cur_steps = [None for _ in range(16)]
cur_len = 0
for i in range(4):
    for j in range(4):
        if grid[i][j] > 0:
            cur_steps[grid[i][j] - 1] = [i, j]
            cur_len += 1
cur_steps = cur_steps[:cur_len]
# print(cur_steps)
ans = 0
for i in range(len(cur_steps) - 1, -1, -1):
    x, y = cur_steps[i]
    if dfs(x, y, i + 1):
        break
    ans += 1
    grid[x][y] = 0

print(ans)

