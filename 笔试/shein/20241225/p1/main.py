from math import inf
time_list = [[1, 4], [2,5], [7, 9]]

time_list.sort()
time_list.append([inf, inf])
n = len(time_list)
cur_l, cur_r = time_list[0]
i = 1
ans = []
while i<n:
    val = time_list[i]
    if val[0] <= cur_r+1:
        cur_r = max(cur_r, val[1])
    else:
        ans.append([cur_l, cur_r])
        cur_l, cur_r = val
    i += 1

print(ans)
