n, m = map(int, input().split())

arr = list(map(int, input().split()))

# 8 8
# 0 1 2 3 4 5 6 0 1 2 3 4 5 6

arr.sort()
for i in range(0, m):
    arr.append(arr[i] + n)

pre_sums = [0 for _ in range(2 * m + 1)]
for i in range(1, 2 * m + 1):
    pre_sums[i] = pre_sums[i - 1] + arr[i - 1]

# 0 1 2 3 4 5 6 7
#     1
left, right = m // 2, m // 2 + m - 1


def get_dist_sum1(pos: int) -> int:
    left_half = m // 2
    right_half = m // 2
    if m % 2 == 0:
        left_half -= 1
    left_sum = pre_sums[pos] - pre_sums[pos - left_half]
    left_dist = left_half * arr[pos] - left_sum
    right_sum = pre_sums[pos + right_half + 1] - pre_sums[pos + 1]
    right_dist = right_sum - right_half * arr[pos]
    return left_dist + right_dist


min_val, min_no = float('inf'), -1
for i in range(left, right + 1):
    cur = get_dist_sum1(i)
    if cur < min_val:
        min_val = cur
        min_no = arr[i]

print(min_no if min_no <= n else min_no - 4)
