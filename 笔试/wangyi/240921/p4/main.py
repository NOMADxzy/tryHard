
# 有一点出发的n个向量，从这n个向量中选a个，怎么样选可以使这a个向量两两之间的夹角的最小值 最大？

n, m = map(int, input().split())

rs = []
deltas = []
for i in range(n):
    a = float(input())
    rs.append(a)
rs.sort()

for i in range(1, n):
    deltas.append((rs[i] - rs[i-1]))
deltas.append(rs[0] + 360 - rs[-1])

min_r_when_choose = [min(deltas)]
for i in range(n, 1, -1):
    # print(deltas)
    min_idx, min_val, min_sum = -1, 360, 360
    for j in range(0, len(deltas)):
        # cur = deltas[j] + deltas[(j+1)%len(deltas)]
        x, x1 = deltas[j], deltas[(j+1)%len(deltas)]
        cur_val = min(x, x1)
        cur_sum = x + x1
        if cur_val < min_val or cur_val == min_val and cur_sum < min_sum:
            min_idx = j
            min_val = cur_val
            min_sum = cur_sum
    new_deltas = []
    idx = 0 if min_idx<len(deltas)-1 else 1
    # print(min_idx)
    while idx<len(deltas):
        if idx == min_idx:
            new_deltas.append(deltas[idx] + deltas[(idx+1) % len(deltas)])
            idx += 1
        else:
            new_deltas.append(deltas[idx])
        idx += 1
    deltas = new_deltas
    min_r_when_choose.append(min(new_deltas))

min_r_when_choose[-1] = 180.00000
min_r_when_choose.reverse()
# print(min_r_when_choose)
max_d = 0
best_a = 0
best_k = 0
for i in range(min(m, n)):
    a = i+1
    k = min_r_when_choose[a-1]
    cur_d = a / (2*m) + k / (360*m) + 0.5 - 1 / (2*m)
    if cur_d >= max_d:
        best_a = a
        best_k = k
        max_d = cur_d

print(f'{best_k:.5f}')
print(best_a)


