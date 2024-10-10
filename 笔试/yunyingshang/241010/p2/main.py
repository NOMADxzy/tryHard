# 区间[a, b] 内 有多少个数可以被c整除
a, b, c = map(int, input().split())

left = (a // c) * c
if a % c > 0:
    left += c
right = (b // c) * c

print((right-left) // c + 1)