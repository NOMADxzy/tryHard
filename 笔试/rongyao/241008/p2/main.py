# 复数求模
from math import pow
a = int(input())
b = int(input())

ans = pow(a*a + b*b, 0.5)
print(round(ans))