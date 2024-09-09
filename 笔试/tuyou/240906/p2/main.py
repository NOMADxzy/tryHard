
n, k = map(int, input().split())
s = input()

# 7 1
# AABABBA

def valid(t) -> bool:
    tmp_map = {}
    tmp_max_val = 0
    for i in range(t):
        if s[i] not in tmp_map:
            tmp_map[s[i]] = 0
        tmp_map[s[i]] += 1
        tmp_max_val = max(tmp_max_val, tmp_map[s[i]])

    if tmp_max_val + k >= t:
        return True

    for i in range(t, n):
        if s[i] not in tmp_map:
            tmp_map[s[i]] = 0
        tmp_map[s[i]] += 1
        tmp_map[s[i-t]] -= 1
        if tmp_map[s[i]] + k >= t:
            return True
    return False


l, r = 1, n
if valid(n):
    print(n)
else:
    while l < r:
        mid = (l+r+1)//2
        if valid(mid):
            l = mid
        else:
            r = mid-1
    print(l)
