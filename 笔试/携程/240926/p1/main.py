# 问题：arr2 是否是 arr1 的子矩阵？

T = int(input())

for _ in range(T):
    m, n = map(int, input().split())
    arr1 = []
    arr2 = []
    for _ in range(m):
        arr1.append(input())
    for _ in range(n):
        arr2.append(input())

    cur_ok = False
    for i in range(m - n + 1):
        if cur_ok:
            break
        for j in range(m - n + 1):
            if cur_ok:
                break
            ok = True
            for k in range(n):
                if not ok:
                    break
                for l in range(n):
                    if arr1[i + k][j + l] != arr2[k][l]:
                        ok = False
                        break
            if ok:
                cur_ok = True
    print("Yes" if cur_ok else "No")

