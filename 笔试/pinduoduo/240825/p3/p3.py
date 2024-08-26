

n, x = map(int, input().split())
arr = list(map(int, input().split()))

# 1 2 3 4 5
# 6

# if 前面有比4大的，就交换 or 前面无序 且 当前持有的比4大

# 2 1 3 2 5
# 4

# 2 1 3 4 5
# 2

# 2 2 3 4 5
# 1

premaxs, pre_is_sorted = [0 for _ in range(n)], [False for _ in range(n)]
premaxs[0] = arr[0]
pre_is_sorted[0] = True
cnts = 0
for i in range(1, n):
    premaxs[i] = max(premaxs[i-1], arr[i])
    pre_is_sorted[i] = True if pre_is_sorted[i-1] and arr[i] >= arr[i-1] else False

if not pre_is_sorted[n-1]:
    cnts = -1
    for i in range(n-1, -1, -1):
        cur_val = arr[i]
        if i>0 and (premaxs[i-1]>cur_val or not pre_is_sorted[i-1] and x > cur_val):
            if premaxs[i]>x or cur_val>=x:
                break
            arr[i] = x
            x = cur_val
            if cnts == -1:
                cnts = 0
            cnts += 1
        if i<n-1 and arr[i] > arr[i+1]:
            cnts = -1
            break

print(cnts)
