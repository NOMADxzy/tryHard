import heapq

N = int(input())

# (1, 200), (1, 200), (2, 10), (3, 100), (3, 100)

points = []
points_map = {}


def add_points(t, w):
    if t not in points_map:
        points_map[t] = []
        points.append(t)
    points_map[t].append(w)


# tasks = []
for _ in range(N):
    d, p = map(int, input().split())
    # tasks.append((d, p))
    add_points(d, p)

points.sort()
points.reverse()
points.append(0)

# print(points)
# print(points_map)

task_queue = []
ans = 0
for i in range(len(points)):
    t = points[i]
    # print("t---{}".format(t))
    for w in points_map.get(t, []):
        heapq.heappush(task_queue, -w)

    if t == 0:
        break
    cnt = t - points[i + 1]
    # print("cnt:", str(cnt))
    while cnt > 0 and len(task_queue) > 0:
        w = -heapq.heappop(task_queue)
        # print(w)
        ans += w
        cnt -= 1

print(ans)




