
# 3 2
# 1 2 1
# 1 2 2 3
# 1 1 2 2
n, q = map(int, input().split())
arr = list(map(int, input().split()))
sums = [0]
for x in arr:
    sums.append(sums[-1] + x)

for _ in range(q):
    l1, r1, l2, r2 = map(int, input().split())
    l1 -= 1
    r1 -= 1
    l2 -= 1
    r2 -= 1
    ans = sum(arr)
    if r1<l2 or r2<l1:
        ans += sums[r1+1]-sums[l1] + sums[r2+1] - sums[l2]
    else:
        idxs = [l1,r1,l2,r2]
        idxs.sort()
        i1, i2, i3, i4 = idxs
        ans += sums[i2] - sums[i1] + 3*(sums[i3+1]-sums[i2]) + sums[i4+1] - sums[i3+1]

    print(ans)
