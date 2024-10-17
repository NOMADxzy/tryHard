
n, L, r, thres = map(int, input().split())

grid = []
visited = []
for _ in range(n):
    grid.append(list(map(int, input().split())))
    visited.append([False for _ in range(n)])

border = 2*r + 1

rect_x1, rect_y1, rect_x2, rect_y2, rect_sum = 0, 0, min(n-1, r), min(n-1, r), 0
for i in range(rect_x1, rect_x2+1):
    for j in range(rect_y1, rect_y2+1):
        rect_sum += grid[i][j]

pre_sums_h = [[0 for _ in range(n+1)] for _ in range(n+1)]
pre_sums_v = [[0 for _ in range(n+1)] for _ in range(n+1)]
for i in range(n):
    for j in range(n):
        pre_sums_h[i+1][j+1] = pre_sums_h[i+1][j] + grid[i][j]
        pre_sums_v[i+1][j+1] = pre_sums_v[i][j+1] + grid[i][j]

def rect_cnt():
    return (rect_x2 - rect_x1 + 1) * (rect_y2 - rect_y1 + 1)

mid_x, mid_y = 0, 0
total = n*n
step = 0
ans = [[0 for _ in range(n)] for _ in range(n)]
dirs = [[0,1],[1,0],[0,-1],[-1,0]]
dir_idx = 0

def update(x1, y1, x2, y2, dir,mid_x, mid_y):
    ret = 0
    if dir[0] != 0:
        if dir[0]==1:
            if mid_x-x1==r:
                ret -= pre_sums_h[x1+1][y2+1] - pre_sums_h[x1+1][y1]
                x1 += 1
            x2 = x2+1
            if x2>=n:
                x2 = n-1
            else:
                ret += pre_sums_h[x2+1][y2+1] - pre_sums_h[x2+1][y1]
        else:
            if x2-mid_x==r:
                ret -= pre_sums_h[x2+1][y2+1] - pre_sums_h[x2+1][y1]
                x2 -= 1
            x1 -= 1
            if x1<0:
                x1 = 0
            else:
                ret += pre_sums_h[x1+1][y2+1] - pre_sums_h[x1+1][y1]
    else:
        if dir[1]==1:
            if mid_y-y1==r:
                ret -= pre_sums_v[x2+1][y1+1] - pre_sums_v[x1][y1+1]
                y1 += 1
            y2 += 1
            if y2>=n:
                y2 = n-1
            else:
                ret += pre_sums_v[x2+1][y2+1] - pre_sums_v[x1][y2+1]
        else:
            if y2 - mid_y == r:
                ret -= pre_sums_v[x2 + 1][y2 + 1] - pre_sums_v[x1][y2 + 1]
                y2 -= 1
            y1 -= 1
            if y1 < 0:
                y1 = 0
            else:
                ret += pre_sums_v[x2 + 1][y1 + 1] - pre_sums_v[x1][y1 + 1]

    return x1, y1, x2, y2, ret


ans_cnt = 0
while step<total:
    ans[mid_x][mid_y] = rect_sum
    visited[mid_x][mid_y] = True
    if rect_sum <= rect_cnt() * thres:
        ans_cnt += 1
    step += 1
    if step == total:
        break

    dx, dy = dirs[dir_idx]
    if mid_x+dx < 0 or mid_x+dx == n or mid_y+dy < 0 or mid_y+dy == n or visited[mid_x+dx][mid_y+dy]:
        dir_idx = (dir_idx + 1) % 4

    rect_x1, rect_y1, rect_x2, rect_y2, delta = update(rect_x1, rect_y1, rect_x2, rect_y2, dirs[dir_idx], mid_x, mid_y)
    mid_x += dirs[dir_idx][0]
    mid_y += dirs[dir_idx][1]
    rect_sum += delta

# print("\n".join(map(str, ans)))
print(ans_cnt)
