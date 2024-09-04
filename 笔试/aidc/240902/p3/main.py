from copy import deepcopy

n, m = map(int, input().split())
arr = []

for _ in range(m):
    arr.append(list(map(int, input().split())))


dp = deepcopy(arr[0])
left_max, right_max = deepcopy(dp), deepcopy(dp)
for j in range(1, n-1):
    j_ = n-1-j
    left_max[j] = max(left_max[j-1], left_max[j])
    right_max[j_] = max(right_max[j_+1], right_max[j_])

for i in range(1, m):
    newDp = []
    for j in range(0, n):
        pre_max = 0
        if j-2>=0 and left_max[j-2]>pre_max:
            pre_max = left_max[j-2]
        if j+2<n and right_max[j+2]>pre_max:
            pre_max = right_max[j+2]
        # print(pre_max)
        newDp.append(arr[i][j] + pre_max)
    dp = newDp
    left_max, right_max = deepcopy(dp), deepcopy(dp)
    for j in range(1, n-1):
        j_ = n-1-j
        left_max[j] = max(left_max[j-1], left_max[j])
        right_max[j_] = max(right_max[j_+1], right_max[j_])

print(max(dp))
