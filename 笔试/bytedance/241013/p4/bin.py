n = int(input())
s = input()


# metas = {'a', 'e', 'i', 'o', 'u'}

def is_meta(x):
    if x == 'a' or x == 'e' or x == 'i' or x == 'o' or x == 'u':
        return True
    else:
        return False


left, right = 1, n
dp = {}


def dfs(l, r):
    if r <= l:
        return True
    key = l * n + r
    if key in dp:
        return dp[key]
    ret = False
    if not is_meta(s[l]) and not is_meta(s[r]) or s[l] == s[r]:
        ret = dfs(l + 1, r - 1)
    dp[key] = ret
    return ret


def check(d):
    for l in range(n - d):
        if dfs(l, l + d):
            return True
    return False


while left < right:
    mid = (left + right + 1) // 2
    if check(mid - 1):
        left = mid
    else:
        right = mid - 1

print(left)
