from copy import deepcopy

n = int(input())
s = input()

# 2323322333
dp = [0 for i in range(n + 1)]

for i in range(0, n):
    max_val = 0
    cnts = [0 for _ in range(10)]
    for j in range(i, -1, -1):
        cnts[int(s[j])] += 1
        if j>0 and dp[j]==0:
            continue
        tmp_cnts = deepcopy(cnts)
        tmp_cnts.sort()
        first_not_0 = 0
        while tmp_cnts[first_not_0] == 0:
            first_not_0 += 1
        tmp_cnts = tmp_cnts[first_not_0:]
        b = tmp_cnts[0]
        if b == 1:
            continue
        ok = True
        for v in tmp_cnts:
            if v % b != 0:
                ok = False
                break
        if ok:
            max_val = max(dp[j] + 1, max_val)
    dp[i + 1] = max_val

print(dp[n])

