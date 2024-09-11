
T = int(input())




for _ in range(T):
    a, b = input().split()
    # if len(a) > len(b):
    #     a, b = b, a
    m, n = len(a), len(b)
    dp = [[0 for _ in range( n +1)] for _ in range( m +1)]
    for i in range( m +1):
        dp[i][0] = i
    for j in range( n +1):
        dp[0][j] = j

    for i in range(1, m+ 1):
        for j in range(1, n + 1):
            min_val = dp[i - 1][j - 1] + (1 if a[i - 1] != b[j - 1] else 0)
            min_val = min(min_val, dp[i - 1][j] + 1)
            min_val = min(min_val, dp[i][j - 1] + 1)
            dp[i][j] = min_val

    print(dp[m][n])


