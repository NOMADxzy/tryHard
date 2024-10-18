

# 6
# 0 0
# 1 0
# 1 1
# 3 1
# 5 0
# 7 1

m = int(input())
details = []
pre_sums = [0 for _ in range(m+1)]
for i in range(m):
    y, r = map(int, input().split())
    details.append((y, r))

details.sort()
for i in range(m):
    pre_sums[i + 1] = pre_sums[i] + details[i][1]

pre_val = -1
cur_max_cnt = 0
ans = -1
for i in range(m):
    y = details[i][0]
    if y == pre_val:
        continue
    left_ok_cnt = i - pre_sums[i]
    right_ok_cnt = pre_sums[m] - pre_sums[i]
    # print(left_ok_cnt + right_ok_cnt)
    if left_ok_cnt + right_ok_cnt >= cur_max_cnt:
        cur_max_cnt = left_ok_cnt + right_ok_cnt
        ans = y
    pre_val = y

print(ans)