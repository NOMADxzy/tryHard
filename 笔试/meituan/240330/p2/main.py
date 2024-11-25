n = int(input())
arr = list(map(int, input().split()))
# 5
# 1 3 2 5 4

# 5 6 5 10 8

if n == 1:
    ans = [2 * arr[0]]
elif n == 2:
    ans = [max(2 * arr[0], arr[1]), max(arr[0], 2 * arr[1])]
else:
    max_left, max_right = [0]*n, [0]*n
    max_left[0] = arr[0]
    max_right[-1] = arr[-1]

    for i in range(1, n):
        j = n - 1 - i
        max_left[i] = max(arr[i], max_left[i - 1])
        max_right[j] = max(arr[j], max_right[j + 1])

    ans = [max(max_right[1], 2 * arr[0])]
    for i in range(1, n - 1):
        ans.append(max(max_left[i - 1], max_right[i + 1], 2 * arr[i]))

    ans.append(max(max_left[-2], 2 * arr[-1]))

print(ans)
