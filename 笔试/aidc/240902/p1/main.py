n, u, v = map(int, input().split())
# 13 5 10
ans = 0

while u < n or v < n:
    if u > v:
        u, v = v, u

    f = u
    remain = u
    while remain > 0:
        f += remain % 10
        remain /= 10
    u = f
    ans += 1

print(ans)
