
n = int(input())
arr = list(map(int, input().split()))

# 14
# 5 1 20 10 10 10 10 15 10 20 1 5 10 15

idx_map = {}
x_list = []

def add_point(x: int, idx: int):
    x += 1
    if x not in idx_map:
        idx_map[x] = []
        x_list.append(x)
    idx_map[x].append(idx)

cur_cnt = 0
pre_val = -1
for i in range(n):
    x = arr[i]
    if x == 0:
        if pre_val > 0:
            cur_cnt += 1
    else:
        add_point(x, i)
        if i == n-1:
            cur_cnt += 1
    pre_val = x

ans = cur_cnt

x_list.sort()
for x in x_list:
    for idx in idx_map[x]:
        arr[idx] = 0
        if idx-1>=0 and idx+1<n and arr[idx-1] > 0 and arr[idx+1] > 0:
            cur_cnt += 1
        elif (idx-1<0 or arr[idx-1]==0) and (idx+1>=n or arr[idx+1]==0):
            cur_cnt -= 1
    ans = max(ans, cur_cnt)

print(ans)