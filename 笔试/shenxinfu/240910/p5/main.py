
n = int(input())
arr = list(map(int, input().split()))
target = int(input())

# 2 1 4 5
# 11
arr.sort()
dp = [False for _ in range(target+1)]
dp[0] = True

for i in range(1, target+1):
    cur = False
    for j in range(0, len(arr)):
        if arr[j] > target:
            break
        if dp[i - arr[j]]:
            cur = True
            break

    dp[i] = cur

print(dp[target])
