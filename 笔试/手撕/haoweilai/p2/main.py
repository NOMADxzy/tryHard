# 对一个字符串s，从0索引开始第奇数个k长度区间的字符进行反转

s = input()
k = int(input())
# abcd
# 2

s_arr = []
for c in s:
    s_arr.append(c)
i = k


def reverse(l, r):
    for i in range((r - l) // 2):
        s_arr[l + i], s_arr[r - 1 - i] = s_arr[r - 1 - i], s_arr[l + i]


while i <= len(s_arr):
    reverse(i - k, i)
    i += k
    if i >= len(s_arr):
        break
    elif i + k >= len(s_arr):
        reverse(i, len(s_arr))
    i += k

print(''.join(s_arr))