from copy import deepcopy
# T = int(input())
T = 1

# 10 4 1
# 2 3 5 7

for _ in range(T):
    target, n, c = map(int, input().split())
    arr = list(map(int, input().split()))
    arr.sort()

    dp = {0}
    ok = c >= target
    for i in range(n):
        if ok:
            break
        newDp = set(dp)
        newDp.add(0)
        for x in dp:
            new_x = x + arr[i]
            if new_x > target:
                continue
            if new_x + c >= target:
                ok = True
                break
            newDp.add(new_x)

        dp = newDp

    print("YES" if ok else "NO")

