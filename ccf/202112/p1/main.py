n, N = map(int, input().split())

arr = list(map(int, input().split()))

acc = 0
pre_val = 0


for i in range(n):
    acc += i * (arr[i] - pre_val)
    pre_val = arr[i]

if N>n:
    acc += (N - arr[-1]) * n

print(acc)