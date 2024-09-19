
# 一个数组每个元素都大于0，每次可以选一个数x，将其变为pow(x,0.5)向下取整，想操作任意次把数组变为不严格单增的数组
# 求最小的操作次数

from math import floor, pow
n = int(input())

arr = list(map(int, input().split()))

# 3
# 2 6 3
ans = 0
for i in range(n-2, -1, -1):
    x = arr[i]
    while x > arr[i+1]:
        x = floor(pow(x, 0.5))
        ans += 1
    arr[i] = x

print(ans)