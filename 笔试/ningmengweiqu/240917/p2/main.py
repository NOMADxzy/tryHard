
# 给一个字符串 和 一个pattern，'.'表示匹配1个字符，'*'表示匹配0个或任意个字符
# 问二者能否成功匹配

s = input()
pattern = input()

m,n = len(s), len(pattern)

dp = [[False for _ in range(n+1)] for _ in range(m+1) ]

dp[0][0] = True

for i in range(1,m+1):
    for j in range(1, n+1):
        if pattern[j-1] == '.':
            dp[i][j] = dp[i-1][j-1]
        elif pattern[j-1] == '*':
            c = pattern[j-2]
            if c == '.':
                c = s[i-1]
            ok = dp[i][j-2]
            for k in range(i-1, -1, -1):
                if s[k] != c or ok:
                    break
                if dp[k][j-2]:
                    ok = True

            dp[i][j] = ok
        else:
            dp[i][j] = s[i-1] == pattern[j-1] and dp[i-1][j-1]
# print(dp)
print('true' if dp[m][n] else 'false')
