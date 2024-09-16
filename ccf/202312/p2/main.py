
q = int(input())

for _ in range(q):
    n, k = map(int, input().split())
    ans = 1

    x = 2
    while n>1:
        cnt = 0
        mul = 1
        while n%x == 0:
            n /= x
            cnt += 1
            mul *= x
        if cnt>=k:
            ans *= mul
        x += 1

    print(ans)