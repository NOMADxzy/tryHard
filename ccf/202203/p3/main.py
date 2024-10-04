import heapq
from copy import deepcopy

n, m = map(int, input().split())

area = {}
witch_area = []
total_nodes = [i for i in range(n)]
for i, a in enumerate(map(int, input().split())):
    if a not in area:
        area[a] = []
    area[a].append(i)
    witch_area.append(a)

task_execute_areas = {}
task_cnt_of_node = [0 for _ in range(n)]
node_execute_tasks = [set() for _ in range(n)]
g = int(input())

for _ in range(g):
    f, id, with_area, with_id, not_with_id, strict = map(int, input().split())

    if id not in task_execute_areas:
        task_execute_areas[id] = set()

    ans = []
    candidates = []
    if with_area == 0:
        nodes_of_area = total_nodes
    else:
        nodes_of_area = area.get(with_area, [])
    if with_id==0:
        candidates = nodes_of_area
    else:
        for node_idx in nodes_of_area:
            if witch_area[node_idx] in task_execute_areas[with_id]:
                candidates.append(node_idx)
    new_candidates = []
    if not_with_id == 0:
        new_candidates = candidates
    else:
        for node_idx in candidates:
            if not_with_id not in node_execute_tasks[node_idx]:
                new_candidates.append(node_idx)
        if strict == 0 and len(new_candidates) == 0:
            new_candidates = candidates
    if len(new_candidates) == 0:
        print(" ".join(["0" for _ in range(f)]))
        continue
    h_arr = [(task_cnt_of_node[node_id], node_id) for node_id in new_candidates]
    heapq.heapify(h_arr)
    for i in range(f):
        if len(h_arr)==0:
            break
        e = heapq.heappop(h_arr)
        task_cnt_of_node[e[1]] += 1
        ans.append(e[1] + 1)
        node_execute_tasks[e[1]].add(id)
        task_execute_areas[id].add(witch_area[e[1]])
        if id == not_with_id:
                continue
        heapq.heappush(h_arr, (task_cnt_of_node[e[1]], e[1]))
    if len(ans) < f:
        if strict == 0:
            h_arr = [(task_cnt_of_node[node_id], node_id) for node_id in new_candidates]
            heapq.heapify(h_arr)
            while len(ans)<f:
                e = heapq.heappop(h_arr)
                task_cnt_of_node[e[1]] += 1
                ans.append(e[1] + 1)
                node_execute_tasks[e[1]].add(id)
                task_execute_areas[id].add(witch_area[e[1]])
                heapq.heappush(h_arr, (task_cnt_of_node[e[1]], e[1]))
        else:
            while len(ans)<f:
                ans.append(0)
    print(" ".join(map(str, ans)))