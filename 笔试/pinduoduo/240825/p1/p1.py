
# 2
# 3
# 3 3 4
# 1 2 1
# 2 3 2
# 3
# 1 2 5
# 1 2 1
# 2 3 2
T = int(input())
# T = 1

for _ in range(T):
    n = int(input())
    v = list(map(int, input().split()))
    edge_weights = []
    remain_weights = 0
    for i in range(n-1):
        a,b,w = map(int, input().split())
        edge_weights.append(w)
        remain_weights += w
    edge_weights.sort()
    ans = v[0] + remain_weights
    for i in range(n-1):
        remain_weights -= edge_weights[i]
        ans = max(ans, v[i+1] + remain_weights)

    print(ans)
