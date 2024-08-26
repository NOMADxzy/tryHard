from math import inf
T = int(input())
# T = 1

# ou + ou
# ou + ji

# 1 2 3 4 5

# 8 8 4 2
# 3 3 2 1

# 2 8

for _ in range(T):
    n = int(input())
    arr = list(map(int, input().split()))

    min_first_opt = inf
    ou_cnts = 0
    for v in arr:
        cnt = 0
        tmp = v
        while tmp % 2 == 0:
            cnt += 1
            tmp /= 2
        min_first_opt = min(min_first_opt, cnt)
        if cnt>0:
            ou_cnts += 1

    ans = min_first_opt
    if ans>0:
        ou_cnts -= 1
    ans += ou_cnts
    print(ans)
