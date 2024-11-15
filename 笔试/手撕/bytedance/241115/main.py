# 用arr中的数组成不大于limit的最大数是多少

# 1 5 8
# 398450549

arr = list(map(int, input().split()))
limit = int(input())

w = 1
mask = 1
while mask * 10 <= limit:
    mask *= 10
    w += 1

arr.sort(reverse=True)


def dfs(preLimit: bool, mask: int) -> int:
    if mask == 0:
        return 0
    if preLimit:
        i = 0
        x = (limit // mask) % 10
        while arr[i] > x:
            i += 1
        while i >= 0:
            next_ret = dfs(arr[i] == x, mask // 10)
            if next_ret >= 0:
                return arr[i] * mask + next_ret
        return -1
    else:
        return arr[0] * mask + dfs(False, mask // 10)


print(dfs(True, mask))
