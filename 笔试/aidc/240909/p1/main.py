
n = int(input())
MOD = 1000000007

pairs = []
max_h = 0

for _ in range(n):
    l, h = map(int, input().split())
    pairs.append((l, h))
    max_h = max(max_h, h)

cnt1 = [0 for _ in range(max_h+1)]
cntv_defer = [0 for _ in range(max_h+1)]
cntv = [0 for _ in range(max_h+1)]

for pair in pairs:
    l, h = pair
    cnt1[l] += 1
    if h>l:
        cnt1[h] += 1
    if h>l+1:
        cntv_defer[l+1] += 1
        cntv_defer[h] -= 1

acc = 0
for i in range(max_h+1):
    acc += cntv_defer[i]
    cntv[i] = acc

# print(cnt1)
# print(cntv_defer)
# print(cntv)

ans = 0
mask = 1
for i in range(max_h+1):
    if (cntv[i] == 0) and (cnt1[i] % 2) == 0:
        pass
    else:
        # print(i, cntv[i], cnt1[i])
        ans = (ans + mask % MOD) % MOD
    mask *= 2

print(ans)