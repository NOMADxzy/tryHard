# 用a、b、c表示模型中的一个三角形，问任意两个点是哪种类型：1-相邻、2-连通、3-不连通

N, m, T = map(int, input().split())
edge_set = set()

parents = [i for i in range(N)]

def find(i: int) -> int:
    if parents[i] != i:
        parents[i] = find(parents[i])
    return parents[i]


def union(i1, i2):
    parents[find(i1)] = find(i2)


for _ in range(m):
    a,b,c = map(int, input().split())
    key1, key2, key3 = N*a + b, N*a + c, N*b + c
    edge_set.add(key1)
    edge_set.add(key2)
    edge_set.add(key3)

    root1, root2, root3 = find(a), find(b), find(c)
    if root2 != root1:
        union(b, a)
    if root3 != root1:
        union(c, a)

for _ in range(T):
    a, b = map(int, input().split())
    if a*N + b in edge_set or b*N + a in edge_set:
        print(1)
    elif find(a) == find(b):
        print(2)
    else:
        print(3)
