
n, m, k = map(int, input().split())
detail_list = []

max_t = 0
for _ in range(n):
    t, c = map(int, input().split())
    detail_list.append([t, c])
    max_t = max(max_t, t)


l, r = k, max_t
def check(target: int) -> bool:
    cost = 0
    for t, c in detail_list:
        if target > t:
            continue
        cost += (t - target) * c
        if cost > m:
            return False
    return True

if check(l):
    print(l)
else:
    while l < r:
        mid = (l + r) // 2
        if check(mid):
            r = mid
        else:
            l = mid + 1
    print(r)