# 给n个点，m个路径，找出s到t的最短路径
import heapq

input_arr = list(map(int, input().split()))

n, m = input_arr[0], input_arr[1]
# 5 6
# 1 2 2
# 1 3 5
# 2 3 4
# 2 4 7
# 3 4 1
# 3 5 3
# 1 5

nexts = [[] for _ in range(n)]
# grid = [[float('inf')] * n for _ in range(n)]

for arr_idx in range(2, 3*m+2, 3):
    a, b, w = input_arr[arr_idx:arr_idx+3]
    a -= 1
    b -= 1

    # grid[a][b] = min(grid[a][b], w)
    # grid[b][a] = min(grid[b][a], w)
    nexts[a].append([b,w])
    nexts[b].append([a,w])

s, t = input_arr[-2:]
s -= 1
t -= 1
dists = [[float('inf'), i] for i in range(n)]
dist_map = [float('inf') for _ in range(n)]
visited = [False] * n
dists[s][0] = 0
dist_map[s] = 0
heapq.heapify(dists)

while True:
    e = heapq.heappop(dists)
    cur_dist, cur_idx = e
    if visited[cur_idx]:
        continue
    if cur_dist == float('inf'):
        break
    if cur_idx == t:
        break

    visited[cur_idx] = True

    for np, nw in nexts[cur_idx]:
        if not visited[np] and nw + cur_dist < dist_map[np]:
            dist_map[np] = nw+cur_dist
            heapq.heappush(dists, [nw+cur_dist, np])

print(dist_map[t] if dist_map[t] < float('inf') else "无法到达")

