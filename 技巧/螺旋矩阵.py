n = int(input())

ans = [[0] * n for _ in range(n)]
idx = 1
dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]]
dir_idx = 0

i, j = 0, 0
while idx <= n * n:
    ans[i][j] = idx

    dx, dy = dirs[dir_idx]
    ni, nj = i + dx, j + dy
    if ni < 0 or ni >= n or nj < 0 or nj >= n or ans[ni][nj] > 0:
        dir_idx = (dir_idx + 1) % 4
        dx, dy = dirs[dir_idx]
        ni, nj = i + dx, j + dy
    i, j = ni, nj
    idx += 1

print("\n".join([" ".join(map(str, row)) for row in ans]))

# "".join(map(str, ans))
