s = input()
# 找出s中最长的一对子字符串，二者的0和1数量相等
# 11011
n = len(s)
left, right = 0, n - 1
start_vals = [0] * n
start_vals[0] = 1 if s[0] == "1" else -1
for i in range(1, n):
    if s[i] == "1":
        start_vals[i] = start_vals[i - 1] + 1
    else:
        start_vals[i] = start_vals[i - 1] - 1


# print(start_vals)

def check(length: int):
    cur_val = start_vals[length - 1]
    l, r = 0, length - 1
    existed_left = {cur_val: 0}
    while r + 1 < n:
        r += 1
        add_val = 1 if s[r] == "1" else -1
        del_val = 1 if s[l] == "1" else -1
        cur_val = cur_val + add_val - del_val
        l += 1
        if cur_val not in existed_left:
            existed_left[cur_val] = l
        else:
            return [existed_left[cur_val] + 1, existed_left[cur_val] + length, l + 1, r + 1]
    return []


while left < right:
    mid = (left + right + 1) // 2
    if check(mid):
        left = mid
    else:
        right = mid - 1

ans = check(left)
print(" ".join(map(str, ans)))
