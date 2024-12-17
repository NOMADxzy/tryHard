# 找数组峰值

arr = [1, 2, 3, 4, 1]
n = len(arr)


def check(pos):
    x = arr[pos]
    l_val, r_val = -float('inf'), -float('inf')
    if pos > 0:
        l_val = arr[pos - 1]
    if pos < n - 1:
        r_val = arr[pos + 1]

    if x > l_val and x > r_val:
        return 0
    elif l_val < x < r_val:
        return -1
    else:
        return 1


left, right = 0, n - 1
while left < right:
    mid = (left + right + 1) // 2
    if check(mid) == 0:
        left, right = mid, mid
        break
    elif check(mid) == -1:
        left = mid
    else:
        right = mid - 1

print(left)
