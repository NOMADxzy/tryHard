from math import sin, cos, pi

n, m = map(int, input().split())
opts = []
for _ in range(n):
    opts.append(tuple(map(float, input().split())))

# 10 5
# 2 0.59
# 2 4.956
# 1 0.997
# 1 1.364
# 1 1.242
# 1 0.82
# 2 2.824
# 1 0.716
# 2 0.178
# 2 4.094
# 1 6 -953188 -946637
# 1 9 969538 848081
# 4 7 -114758 522223
# 1 9 -535079 601597
# 8 8 159430 -511187

sums_k = [0 for _ in range(n+1)]
sums_theta = [0 for _ in range(n+1)]
sums_k[0] = 1
for i in range(n):
    opt, val = opts[i]
    sums_k[i+1], sums_theta[i+1] = sums_k[i], sums_theta[i]
    if opt == 1:
        sums_k[i+1] = sums_k[i] * val
    elif opt == 2:
        sums_theta[i+1] = sums_theta[i] + val

for _ in range(m):
    i, j, x, y = map(int, input().split())
    k, theta = sums_k[j] / sums_k[i-1], (sums_theta[j] - sums_theta[i-1]) % (2*pi)

    x, y = k*x, k*y
    x, y = x*cos(theta) - y*sin(theta), x*sin(theta) + y*cos(theta)
    print(round(x, 3), round(y, 3))

