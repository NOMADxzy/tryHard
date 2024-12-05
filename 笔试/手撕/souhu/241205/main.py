# 给定一个数组，和min_val, max_val，求最大值在[min_val, max_val]之间的子数组数量

# 15 4 5 9 1 14 3 14 10
# 9 10

arr = list(map(int, input().split()))
min_val, max_val = map(int, input().split())
n = len(arr)

largers = [-1]
for i in range(n):
    if arr[i]>max_val:
        largers.append(i)

ans = 0
largers.append(n)
for i in range(len(largers)-1):
    sub_arr = arr[largers[i]+1:largers[i+1]]
    if len(sub_arr) == 0:
        continue
    total_cnt = 0
    pre_cnt = 0
    m = len(sub_arr)
    for j in range(m):
        cur_cnt = 0
        if sub_arr[j]>=min_val:
            cur_cnt = (j+1)
        else:
            cur_cnt = pre_cnt
        total_cnt += cur_cnt
        pre_cnt = cur_cnt
    ans += total_cnt

print(ans)
