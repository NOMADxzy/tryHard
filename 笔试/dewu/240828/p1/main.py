target = input()
p, q = map(int, input().split())

n = len(target)
# bbcabc
# 3 1

dp = [[0] * n for _ in range(n)]
for i in range(n):
    dp[i][i] = p


def is_sub(a, b) -> (bool, bool):
    left_main = True
    if len(a) < len(b):
        left_main = False
        a, b = b, a
    for i in range(len(a)):
        if i + len(b) > len(a):
            break
        ok = True
        for j in range(len(b)):
            if a[i + j] != b[j]:
                ok = False
                break
        if ok:
            return True, left_main
    return False, False


for l in range(1, n):
    for i in range(0, n - l):
        left, right = i, i + l
        min_val = min(dp[left][right - 1] + p, dp[left + 1][right] + p)
        for k in range(left, right):
            sub, left_main = is_sub(target[left:k + 1], target[k + 1:right + 1])
            if sub:
                if left_main:
                    min_val = min(min_val, dp[left][k] + q)
                else:
                    min_val = min(min_val, dp[k + 1][right] + q)
        dp[left][right] = min_val

print(dp[0][n - 1])

31331
