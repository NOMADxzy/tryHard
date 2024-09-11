T = int(input())

# 6 3
# 8 1 10 1 1 1
for _ in range(T):
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    idxs = [i for i in range(n)]
    idxs.sort(key=lambda x: arr[x])
    idxs = idxs[:k]
    idxs_map = {}
    for i in idxs:
        idxs_map[i] = True
    new_tasks = []
    total = 0
    for i in range(n):
        if i in idxs_map:
            new_tasks.append(arr[i])
            total += arr[i]
    # print(idxs)
    # print(new_tasks)
    acc = 0
    i = 0
    while acc * 2 < total:
        acc += new_tasks[i]
        i += 1
    # print(total, acc)
    ans = min(acc, total - (acc - new_tasks[i - 1]))
    print(ans)


