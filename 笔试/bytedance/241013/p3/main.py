# MEX(arr)表示 arr中未出现的最小非负整数，将数组划分若干个子数组，各个部分的MEX值之和最大为多少
T = int(input())

# 0 1 1 0

def dfs(arr, pos, dp):
    if pos == len(arr):
        return 0
    if dp[pos]>=0:
        return dp[pos]
    exsits = set()
    target = 0
    max_val = 0
    for right in range(pos, len(arr)):
        exsits.add(arr[right])
        if arr[right] == target:
            target += 1
            while target in exsits:
                target += 1
            max_val = max(max_val, target+dfs(arr, right+1, dp))
    dp[pos] = max_val
    return max_val


for _ in range(T):
    n = int(input())
    arr = list(map(int, input().split()))
    ans = dfs(arr, 0, [-1]*n)
    print(ans)
