# 从n个点开始，广度优先的着色，最终那些着色的点中比所有邻居都晚（或相等）着色的点称为重要点
# 输出所有的重要点

n, m, t = map(int, input().split())
nexts = [[] for _ in range(n+1)]

for _ in range(m):
    u, v = map(int, input().split())
    nexts[u].append(v)
    nexts[v].append(u)

queue = list(map(int, input().split()))
cur_time = 1
visited = {}
for e in queue:
    visited[e] = 0
ans = []
while len(queue)>0:
    new_queue = []
    # re_check = deepcopy(queue)
    while len(queue)>0:
        e = queue[0]
        queue = queue[1:]

        for np in nexts[e]:
            if np not in visited:
                visited[np] = cur_time
                new_queue.append(np)
    queue = new_queue
    cur_time += 1
    # print(queue)

for e in range(1, n+1):
    is_best = e in visited
    if not is_best:
        continue
    for np in nexts[e]:
        if visited[np] > visited[e]:
            is_best = False
            break
    if is_best:
        ans.append(e)

ans.sort()
print(" ".join(map(str, ans)))

