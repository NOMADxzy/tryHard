
n, m = map(int, input().split())

matrix = [{} for _ in range(n)]
best_obj = [[0, -1] for _ in range(n)]

def cpt_best(node_idx):
    best_val, best_idx = 0, -1
    del_keys = []
    for obj, val in matrix[node_idx].items():
        if val == 0:
            del_keys.append(obj)
            continue
        if val > best_val or val == best_val and obj < best_idx:
            best_idx, best_val = obj, val
    for k in del_keys:
        matrix[node_idx].pop(k)
    best_obj[node_idx][0], best_obj[node_idx][1] = best_val, best_idx

def update(u, v, delta):
    tmp_dict = matrix[u]
    if v not in tmp_dict:
        tmp_dict[v] = 0
    new_val = tmp_dict[v] + delta
    tmp_dict[v] = new_val
    if new_val > best_obj[u][0] or new_val == best_obj[u][0] and v < best_obj[u][1]:
        best_obj[u][0], best_obj[u][1] = new_val, v

release_map = {}

for cur_day in range(m):
    tmp = release_map.pop(cur_day, [])
    for u,v,x in tmp:
        matrix[u][v] -= x
        if v == best_obj[u][1]:
            cpt_best(u)
        matrix[v][u] -= x
        if u == best_obj[v][1]:
            cpt_best(v)



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
        main_object = best_obj[node_idx][1]
        print(main_object + 1)

    do_query_island, do_query_pair = map(lambda x: bool(int(x)), input().split())
    if do_query_island or do_query_pair:
        island_cnt, pair_cnt = 0, 0
        main_objs = [best_obj[node_idx][1] for node_idx in range(n)]
        for node_idx in range(n):
            if do_query_pair and main_objs[node_idx] > node_idx and main_objs[main_objs[node_idx]] == node_idx:
                pair_cnt += 1
            if do_query_island and main_objs[node_idx] == -1:
                island_cnt += 1
        if do_query_island:
            print(island_cnt)
        if do_query_pair:
            print(pair_cnt)




