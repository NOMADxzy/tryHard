import heapq

n, m = map(int, input().split())

matrix = [{} for _ in range(n)]
heaps = [[] for _ in range(n)]
# best_obj = [[0, -1] for _ in range(n)]
main_objs = [-1 for _ in range(n)]
# island_cnt = n


def update(u, v, delta):
    global island_cnt
    state_dict = matrix[u]
    if v not in state_dict:
        state_dict[v] = 0
    state_dict[v] += delta
    new_val = state_dict[v]
    heapq.heappush(heaps[u], (-new_val, v))

    h_arr = heaps[u]
    while len(h_arr) > 0 and -h_arr[0][0] != matrix[u][h_arr[0][1]]:
        heapq.heappop(h_arr)
    if len(h_arr) > 0 and h_arr[0][0] < 0:
        # if main_objs[u] == -1:
        #     island_cnt -= 1
        main_objs[u] = h_arr[0][1]
    else:
        # if main_objs[u] != -1:
        #     island_cnt += 1
        main_objs[u] = -1


release_map = {}

for cur_day in range(m):
    tmp = release_map.pop(cur_day, [])
    for u,v,x in tmp:
        update(u, v, -x)
        update(v, u, -x)


    splits = list(map(int, input().split()))
    for i in range(1, len(splits), 4):
        u,v,x,y = splits[i: i+4]
        u -= 1
        v -= 1
        update(u, v, x)
        update(v, u, x)

        release_day = cur_day + y
        if release_day not in release_map:
            release_map[release_day] = []
        release_map[release_day].append((u, v, x))


    query_nodes = list(map(int, input().split()))[1:]
    for node_idx in query_nodes:
        node_idx -= 1
        print(main_objs[node_idx] + 1)

    do_query_island, do_query_pair = map(lambda x: bool(int(x)), input().split())
    if do_query_island or do_query_pair:
        island_cnt, pair_cnt = 0, 0

        for node_idx in range(n):
            if do_query_pair and main_objs[node_idx] > node_idx and main_objs[main_objs[node_idx]] == node_idx:
                pair_cnt += 1
            if do_query_island and main_objs[node_idx] == -1:
                island_cnt += 1
        if do_query_island:
            print(island_cnt)
        if do_query_pair:
            print(pair_cnt)




