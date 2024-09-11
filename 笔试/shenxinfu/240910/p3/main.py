from math import inf
n, m = map(int, input().split())

idxs = {}

for _ in range(n):
    splits = input().split()
    idxs[splits[0]] = int(splits[1])-1

nexts = [[] for _ in range(n)]

for _ in range(m):
    a, b = map(int, input().split())
    a -= 1
    b -= 1
    nexts[a].append(b)
    nexts[b].append(a)

dp = {}
def dfs(a, b, visited)->int:
    key = a * n + b
    # if key in dp:
    #     return dp[key]
    if a == b:
        return 0
    visited[a] = True
    ret = inf
    for np in nexts[a]:
        if np in visited and visited[np]:
            continue
        ret = min(ret, dfs(np, b, visited) + 1)
    visited[a] = False
    # dp[key] = ret
    return ret

# print(dfs(0, 5, {}))
q = int(input())
for _ in range(q):
    s1, s2 = input().split()
    i1, i2 = idxs[s1], idxs[s2]

    ans = dfs(i1, i2, {})
    if ans == inf:
        print(-1)
    else:
        print(ans)
