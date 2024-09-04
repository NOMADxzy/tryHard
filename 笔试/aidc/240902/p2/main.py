
T = int(input())

for _ in range(T):
    length, n, k = map(int, input())
    obstacle_map = {}
    obstacles = list(map(int, input().split()))
    costs = list(map(int, input().split()))
    for i in len(obstacles):
        obstacle_map[obstacles[i]] = costs[i]

    p = k-1
    cur_time = 0
    while p<length:
        if p in obstacle_map:
            cur_time += obstacle_map[p]
        if cur_time <= p:
            break
        # moments[p] = cur_time
        cur_time += 1
        p += 1
    
    if p==length:
        print("YES")
    else:
        print("NO")