# 将一个整数表示成素数相乘的表达式

n = 12

is_comp = [False] * (n+1)
primes = []

for i in range(2, n+1):
    if is_comp[i]:
        continue
    primes.append(i)
    for x in range(i*i, n+1, i):
        is_comp[x] = True

# 2^2 3^1 5^1
# (2+1) * (1+1) * (1+1)
ans = ""
cur_idx = 0

while n > 1:
    cur = primes[cur_idx]
    cur_cnt = 0
    while n%cur == 0:
        cur_cnt += 1
        n //= cur
    ans += "*" + str(cur) + "^" + str(cur_cnt)
    cur_idx += 1

print(ans[1:])
