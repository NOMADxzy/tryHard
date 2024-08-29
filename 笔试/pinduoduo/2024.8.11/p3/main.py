
# 1 0 1 1 1 0 0 1

# 1 1 0 0 0 1 1 0


# TOTAL - x + (-x)
# TOTAL - 2x
# 0 + 2
# 0 - 2
# 5 - 2 * 2

# 0: 3
# 1: 5

# 0 1 0 0 0 1 1 0

n = int(input())
arr = list(map(int, input().split()))
arr = [1 if num == 1 else -1 for num in arr]
TOTAL = sum(arr)
lefts = [[0,0] for _ in range(n)]
rights = [[0,0] for _ in range(n)]

min_val, max_val = 0,0
acc = 0
for i in range(0, n-1):
    acc += arr[i]
    min_val = min(min_val, acc)
    max_val = max(max_val, acc)
    lefts[i][0] = min_val
    lefts[i][1] = max_val


min_val, max_val = 0,0
acc = 0
for i in range(n-2, -1, -1):
    acc += arr[i+1]
    min_val = min(min_val, acc)
    max_val = max(max_val, acc)
    rights[i][0] = min_val
    rights[i][1] = max_val


total_min_val, total_max_val = 0,0
for i in range(0, n-1):
    total_max_val = max(total_max_val, TOTAL - (lefts[i][0] + rights[i][0]))
    total_min_val = min(total_min_val, TOTAL - (lefts[i][1] + rights[i][1]))

if total_min_val*total_max_val >= 0:
    print(total_max_val - total_min_val + 1)
else:
    total_min_val = -total_min_val
    val = max(total_min_val, total_max_val)
    print(val + 1)



# acc = 0
# min_l, max_l = 0, 0
# min_r, max_r = 0, 0
#
# for num in arr:
#     acc += num
#     min_r = min(min_r, acc - max_l)
#     max_r = max(max_r, acc - min_l)
#
#     min_l = min(min_l, acc)
#     max_l = max(max_l, acc)
#
# min_val, max_val = min_l + min_r, max_l + max_r
#
# if min_val*max_val >= 0:
#     print(max_val - min_val + 1)
# else:
#     min_val = -min_val
#     val = max(min_val, max_val)
#     print(val + 1)