from math import floor

n, N = map(int, input().split())

arr = list(map(int, input().split()))

f_val = 0
r = floor(N / (n+1))
def g_val(x):
    return floor(x / r)

def acc_g_val(left_i, right_i):
    ret = 0

    itl_len = right_i - left_i + 1
    head_need = left_i % r
    tail_need = r - 1 - (right_i % r)
    full_epoch = (itl_len + head_need + tail_need) // r
    start_val = g_val(left_i)
    end_val = g_val(right_i)
    ret += r * (start_val + end_val) * full_epoch // 2
    ret -= start_val * head_need
    ret -= end_val * tail_need

    return ret


pre_x = 0
acc_error = 0

arr.append(N)
n += 1
for i in range(n):
    left = r * i - 1
    right = r * (i + 1)
    if left >= pre_x:
        left = min(left, arr[i]-1)
        acc_f = (left-pre_x+1) * i
        acc_g = acc_g_val(pre_x, left)
        acc_error += abs(acc_g - acc_f)
    if right <= arr[i]-1:
        right = max(pre_x, right)
        acc_f = (arr[i] - right) * i
        acc_g = acc_g_val(right, arr[i]-1)
        acc_error += abs(acc_g - acc_f)

    pre_x = arr[i]

print(acc_error)