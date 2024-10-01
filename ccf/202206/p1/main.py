import math

n = int(input())
arr = list(map(int, input().split()))

mean_val = sum(arr) / n
d_val = 0

for x in arr:
    d_val += (x - mean_val) * (x - mean_val)

d_val /= n
std_val = math.pow(d_val, 0.5)

for i in range(n):
    arr[i] = (arr[i] - mean_val) / std_val

for x in arr:
    print(x)