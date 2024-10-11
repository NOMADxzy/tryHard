from math import inf
q = int(input())
# 3
# 3 3
# 2 5
# 5 3

def sub_interval(t1, t2):
    if t1[1] < t2[0] or t1[0] > t2[1]:
        return None
    else:
        return [max(t1[0], t2[0]), min(t1[1], t2[1])]


cur_interval = [-inf, inf]

for _ in range(q):
    m, d = map(int, input().split())
    cur_interval = sub_interval(cur_interval, [m-d, m+d])
    if cur_interval is None:
        break

print(cur_interval[1] if cur_interval else -1)
