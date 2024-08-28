from math import inf

n, x = map(int, input().split())
arr = list(map(int, input().split()))
arr.sort()
total = sum(arr)

l,r = 0, n-1
ans = inf

while l<r:
    v1 = x*(arr[l] + arr[r])
    v2 = total - arr[l] - arr[r]
    if v1>v2:
        r -= 1
    elif v1<v2:
        l += 1
    else:
        ans = 0
        break

    ans = min(ans, abs(v1 - v2))

print(ans)