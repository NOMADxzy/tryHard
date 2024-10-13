# 求最长的子数组长度，数组内所有元素的最大公约数为1

n = int(input())
arr = list(map(int, input().split()))
# n = len(arr)

l,r = 0,0
ans = 0

top_val = 100000
is_comp = [False]*(top_val+1)
prims = [] # 1. 列表替代集合，因为集合添加元素时开销较大
for x in range(2, ):
    if is_comp[x]:
        continue
    prims.append(x)
    for nx in range(x*x, top_val+1, x):  # 2*x 改为x*x，因为2*x已经被2筛过一次了
        is_comp[nx] = True


last_idx = {}
while r<n:
    x = arr[r]
    cur = 2
    tmp = []
    while x>=cur:
        if x % cur == 0:
            # last_idx[cur] = r
            tmp.append(cur)
            while x >= cur and x % cur == 0:
                x //= cur

        cur += 1

    for x in tmp:
        if x in last_idx and last_idx[x] >= l:
            l = last_idx[x] + 1
        last_idx[x] = r

    ans = max(ans, (r-l+1))
    r += 1

print(ans)
