
n, m = map(int, input().split())

nums1 = list(map(int, input().split()))
nums2 = list(map(int, input().split()))

if n<m:
    n,m = m,n
    nums1, nums2 = nums2, nums1


ans = m + n
# 7 10
# 2 2 1 2 1 1 2
# 2 1 1 2 1 1 2 1 1 2

for i in range(m-1, -1, -1):
    ok = True
    for j in range(i, m):
        cur = nums1[j-i] + nums2[j]
        if cur > 3:
            ok = False
            break
    if ok:
        ans = min(ans, i + n)

for i in range(n):
    ok = True
    for j in range(i, n):
        if j-i>=m:
            break
        cur = nums1[j] + nums2[j-i]
        if cur > 3:
            ok = False
            break
    if ok:
        cur_ans = n
        l1 = n - i
        if l1 < m:
            cur_ans += m - l1
        ans = min(ans, cur_ans)
        break

print(ans)
