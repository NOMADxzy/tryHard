
m, n = map(int, input().split())
dirs = []
dir_map = {"L":0, "R":1, "U":2, "D":3}
dirs = [(0, -1), (0, 1), (-1, 0), (1, 0)]
grid = []

for _ in range(m):
    grid.append([dir_map[c] for c in input()])

dp = [[False for _ in range(n)] for _ in range(m)]
ans = 0

def dfs(p, q, visited)->bool:
    if dp[p][q]:
        return True
    if visited[p][q]:
        # print(p,q)
        return True
    visited[p][q] = True

    dir = dirs[grid[p][q]]
    p += dir[0]
    q += dir[1]

    if p<0 or p>=m or q<0 or q>=n:
        return False

    ret = dfs(p, q, visited)
    if ret:
        dp[p][q] = True
    else:
        dp[p][q] = False
    return ret



for i in range(m):
    for j in range(n):
        if dfs(i, j, [[False for _ in range(n)] for _ in range(m)]):
            ans += 1

print(ans)