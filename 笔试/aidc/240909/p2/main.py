import heapq

n = int(input())
arr = list(map(int, input().split()))

idxs = {}
ans = []
for i, v in enumerate(arr):
    if v not in idxs:
        idxs[v] = []
    same_idxs = idxs[v]
    same_idxs.append(i)
    # if len(same_idxs)>=2:
    #     ans.append(same_idxs[0])
    #     ans.append(same_idxs[1])
    #     same_idxs = same_idxs[2:]

visited = {}
# print(idxs)
for i in range(n):
    if i in visited:
        continue
    else:
        same_idxs = idxs[arr[i]]
        if len(same_idxs) >= 2:
            v1, v2 = same_idxs[0], same_idxs[1]
            # print(v1, v2)
            visited[v1] = True
            visited[v2] = True
            idxs[arr[i]] = same_idxs[2:]
            ans.append(v1 + 1)
            ans.append(v2 + 1)

print(" ".join(map(str, ans)))
