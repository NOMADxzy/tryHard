
n, m = map(int, input().split())
arr = list(map(int, input().split()))
# 10 100
# 82 37 43 121 55 1 1 1 1 1

ans = 0
arr.sort()

left, right = 0, n-1
while left<=right:
    target_cnt = m // arr[right]
    if m % arr[right] > 0:
        target_cnt += 1
    if target_cnt>=2:
        if left+target_cnt-2 < right:
            left += target_cnt-1
            right -= 1
            ans += 1
        else:
            break
    else:
        right -= 1
        ans += 1
print(ans)