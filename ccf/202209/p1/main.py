
n, m = map(int, input().split())
a_list = list(map(int, input().split()))

pres = [1]
for i in range(0, n-1):
    pres.append(pres[-1] * a_list[i])

ans = [-1 for _ in range(n)]
for i in range(n-1, -1, -1):
    ans[i] = m // pres[i]
    m -= pres[i] * ans[i]

print(" ".join(map(str, ans)))