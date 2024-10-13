# 元回文，表示字符串中的元音字母必须满足回文串的条件，求元回文最长子串的长度

n = int(input())
s = input()

metas = {'a', 'e', 'i', 'o', 'u'}

# pre_dp = [True] * n
dp =[True for i in range(n)]

ans1 = 1

for d in range(2, 4, 2):
    newDp = [False] * (n-d)
    for left in range(0, n-d):
        right = left + d
        if min(left, n-1-right)*2 + d+1 <= ans1:
            break
        if (dp[left+1]) and (s[left] not in metas and s[right] not in metas or s[left] == s[right]):
            p, q = left, right
            while p - 1 >= 0 and q + 1 < n and (
                    s[p - 1] not in metas and s[q + 1] not in metas or s[p - 1] == s[q + 1]):
                p -= 1
                q += 1
            ans1 = max(ans1, q - p + 1)
    dp = newDp

dp =[True for i in range(n)]

ans2 = 1

for d in range(1, 3, 2):
    newDp = [False] * (n-d)
    for left in range(0, n-d):
        right = left + d
        if min(left, n-1-right)*2 + d+1 <= ans2:
            break
        if (dp[left+1]) and (s[left] not in metas and s[right] not in metas or s[left] == s[right]):
            p, q = left, right
            while p-1>=0 and q+1<n and (s[p-1] not in metas and s[q+1] not in metas or s[p-1] == s[q+1]):
                p -= 1
                q += 1
            ans2 = max(ans2, q-p+1)
    dp = newDp

print(max(ans1, ans2))