from os import cpu_count

n, m = map(int, input().split())
rewards = list(map(int, input().split()))
b_arr = list(map(int, input().split()))

# 5 20
# 1 2 3 4 5
# 4 2 3 5 10

def cpt_val(l):
    rewards_part = rewards[:l]
    b_arr_part = b_arr[:l]
    b_arr_part.sort()
    rewards_part.sort()
    val = 0
    for i in range(l):
        val += b_arr_part[i] * rewards_part[i]
    return val

tmp = cpt_val(n)
if tmp<m:
    print(-1)
else:
    l, r = 1, n-1
    while l<r:
        mid = (l+r)//2
        if cpt_val(mid) >= m:
            r = mid
        else:
            l = mid + 1
    print(r)