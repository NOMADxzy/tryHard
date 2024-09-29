import heapq

n, m = map(int, input().split())
pres = list(map(int, input().split()))
nexts = [[] for _ in range(m)]
days = list(map(int, input().split()))

start_ans = [1 for _ in range(m)]
end_ans = [0 for _ in range(m)]

cur_nodes = []

for i in range(m):
    pres[i] -= 1
    if pres[i] == -1:
        cur_nodes.append((days[i], i))
    else:
        nexts[pres[i]].append(i)

heapq.heapify(cur_nodes)
cur_time = 0

while len(cur_nodes)>0:
    e = heapq.heappop(cur_nodes)
    cur_time = e[0]
    for new_i in  nexts[e[1]]:
        new_node = (cur_time + days[new_i], new_i)
        heapq.heappush(cur_nodes, new_node)
        start_ans[new_i] = cur_time + 1

print(" ".join(map(str, start_ans)))


if cur_time <= n+1:

    cur_time = n
    cur_nodes = []
    for i in range(m):
        if len(nexts[i]) == 0:
            cur_nodes.append((-(cur_time - days[i]), i))
            end_ans[i] = cur_time - days[i] + 1

    heapq.heapify(cur_nodes)
    while len(cur_nodes) > 0:
        e = heapq.heappop(cur_nodes)
        cur_time = -e[0]
        if pres[e[1]] != -1:
            new_i = pres[e[1]]
            new_node = (-(cur_time - days[new_i]), new_i)
            heapq.heappush(cur_nodes, new_node)
            end_ans[new_i] = cur_time - days[new_i] + 1

    print(" ".join(map(str, end_ans)))