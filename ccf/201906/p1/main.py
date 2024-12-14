
# 3 3
# 73 -8 -6 -4
# 76 -5 -10 -8
# 80 -6 -15 0

n, m = map(int, input().split())
total, max_idx, max_lost_cnt = 0, -1, -1

trees = [0] * n
for i in range(n):
    split_ints = list(map(int, input().split()))
    trees[i] = split_ints[0]
    lost_cnt = -sum(split_ints[1:])
    trees[i] -= lost_cnt
    total += trees[i]
    if lost_cnt > max_lost_cnt:
        max_lost_cnt = lost_cnt
        max_idx = i

print(total, max_idx+1, max_lost_cnt)