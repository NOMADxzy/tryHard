
n = int(input())
arr = list(map(int, input().split()))
# 1 1 4 5 1 4

dp_left, dp_right = [0 for _ in range(n)], [0 for _ in range(n)]


for i in range(1, n):
    j = n - 1 - i

    if arr[i] > arr[i-1]:
        dp_left[i] = dp_left[i-1] + 1

    if arr[j] < arr[j+1]:
        dp_right[i] = dp_right[i+1] + 1


ans = 0
for i in range(1, n-1):
    cur = dp_left[i] + dp_right[i] + 1 if dp_left[i]>0 and dp_right[i]>0 else 0
    ans = max(ans, cur)

print(ans)
