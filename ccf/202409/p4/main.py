from math import inf

# 5 5
# 0 0
# 2 4
# 4 0
# 5 3
# 5 5
# 1 2 2 5
# 3 5 2 6
# 2 0 2 1
# 4 2 2 3
# 5 4 1 2


n, m = map(int, input().split())
points = []
for _ in range(n):
    x, y = map(int, input().split())
    points.append((x, y))

delays = [[inf for _ in range(n-l)] for l in range(n)]

def reachable(x,y,r,x1,y1):
    w = abs(x - x1)
    h = abs(y-y1)
    return w<=r and h<=r

for _ in range(m):
    x, y, r, t = map(int, input().split())
    tmp_list = []
    for i in range(len(points)):
        if reachable(x,y,r,points[i][0], points[i][1]):
            tmp_list.append(i)
    # print(tmp_list)
    for i in range(len(tmp_list)):
        for j in range(i+1, len(tmp_list)):
            a, b = tmp_list[i], tmp_list[j]
            delays[a][b-a] = min(delays[a][b-a], t)

visited = [False for _ in range(n)]
# visited[0] = True
cur_delays = [inf for _ in range(n)]
cur_delays[0] = 0

def get_delay(a, b):
    if a>b:
        a,b = b,a
    return delays[a][b-a]

while True:
    if not visited[0]:
        min_idx, min_val = 0, 0
    else:
        min_idx, min_val = -1, inf
        for i in range(1, n):
            if not visited[i] and cur_delays[i] < min_val:
                min_val = cur_delays[i]
                min_idx = i
        if min_idx == -1:
            break

    visited[min_idx] = True
    if min_idx == n-1:
        break
    for i in range(1, n):
        if not visited[i]:
            new_val = min_val + get_delay(min_idx, i)
            if new_val < cur_delays[i]:
                cur_delays[i] = new_val


print(cur_delays[n-1] if cur_delays[n-1] != inf else "Nan")


