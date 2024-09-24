from functools import cmp_to_key

n = int(input())
arr = list(map(str, input().split()))


def better(a, b):
    v1 = int(a + b)
    v2 = int(b + a)
    if v1 == v2:
        return 0
    else:
        return -1 if v1 > v2 else 1


arr.sort(key=cmp_to_key(better))
if arr[0] == "0":  # 全0的情况，不能输出“0000”
    print("0")
else:
    print("".join(arr))
