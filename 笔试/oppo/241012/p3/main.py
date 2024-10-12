# 给n个数，构造数组，使得相邻数的差绝对值之和最大

n = int(input())
arr = list(map(int, input().split()))
if n < 1:
    ans = arr
else:
    arr.sort(key=lambda x: -x)
    ans = []

    cur = [arr[0], arr[-1]]
    arr = arr[1:-1]
    while len(cur)<n:
        v1, v2 = cur[0], cur[-1]
        if v1 > v2 :
            v1, v2 = v2, v1
        d1 = abs(v1 - arr[0])
        d2 = abs(v2 - arr[-1])
        if d1 >= d2:
            if cur[-1] == v1:
                cur.append(arr[0])
            else:
                cur.insert(0, arr[0])
            arr = arr[1:]
        else:
            if cur[-1] == v2:
                cur.append(arr[-1])
            else:
                cur.insert(0, arr[-1])
            arr = arr[:-1]
    ans = cur


print(" ".join(map(str, ans)))