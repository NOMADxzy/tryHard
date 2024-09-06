
n, k, limit = map(int, input().split())
arr = list(map(int, input().split()))

# 5 3 10
# 9 7 3 6 5

# 9 1 0 5 5

l, r = 0, k-1
# p = r
cur_total = sum(arr[l:r+1])
ans = 0
while r<n:
    if cur_total > limit:
        j = r
        while cur_total > limit:
            if arr[j] > cur_total - limit:
                arr[j] -= cur_total - limit
                ans += cur_total - limit
                cur_total = limit
            else:
                cur_total -= arr[j]
                ans += arr[j]
                arr[j] = 0
            j -= 1
    r += 1
    if r==n:
        break
    cur_total += arr[r]
    cur_total -= arr[l]
    l += 1

print(ans)
