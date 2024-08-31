
a, b, c, k = map(int, input().split())
# 1 2 3 3
first = (a+b+c+k) // 3
remain = (a + b + c + k) - first*3
vals = [first, first, first]

for i in range(len(vals)):
    if remain==0:
        break
    vals[i] += 1
    remain -= 1

MOD = 1000000007
ans = ((vals[0] %MOD) * (vals[1]%MOD)) % MOD
ans = (ans * (vals[2]%MOD)) %MOD

print(ans)

