
a, b = map(int, input().split())


d = b - a

real_b = 0
for i in range(1, d+1):
    real_b += i

print(real_b - b)