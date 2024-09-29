from math import cos, pow, pi

metric_arr = []
for _ in range(8):
    metric_arr.append(list(map(int, input().split())))

n = int(input())
T = int(input())

data = list(map(int, input().split()))

# 16 11 10 16 24 40 51 61
# 12 12 14 19 26 58 60 55
# 14 13 16 24 40 57 69 56
# 14 17 22 29 51 87 80 62
# 18 22 37 56 68 109 103 77
# 24 35 55 64 81 104 113 92
# 49 64 78 87 103 121 120 101
# 72 92 95 98 112 100 103 99
# 26
# 2
# -26 -3 0 -3 -2 -6 2 -4 1 -3 1 1 5 1 2 -1 1 -1 2 0 0 0 0 0 -1 -1

# 填充 -> 量化 -> 余弦变换 + 偏移

arr1 = [[0 for _ in range(8)] for _ in range(8)]
dirs = [[+1, -1], [-1, +1]]
d_idx = 1
p,q, i = 0,0,0
while i<len(data):
    arr1[p][q] = data[i]
    p += dirs[d_idx][0]
    q += dirs[d_idx][1]
    if p < 0 or p > 7 or q < 0 or q > 7:
        p -= dirs[d_idx][0]
        q -= dirs[d_idx][1]
        d_idx = 1 - d_idx
        if p == 0 or p == 7:
            q += 1
        else:
            p += 1
    i += 1

def print_arr(arr):
    for row in arr:
        print(" ".join(map(str, row)))

if T==0:
    print_arr(arr1)

else:
    for i in range(8):
        for j in range(8):
            arr1[i][j] *= metric_arr[i][j]

    if T == 1:
        print_arr(arr1)
    else:
        def alpha(x):
            return pow(0.5, 0.5) if x == 0 else 1
        arr2 = [[0 for _ in range(8)] for _ in range(8)]
        for i in range(8):
            for j in range(8):
                for u in range(8):
                    for v in range(8):
                        arr2[i][j] += alpha(u) * alpha(v) * arr1[u][v] * cos((pi/8) * (i+0.5) * u) * cos((pi/8) * (j+0.5) * v) / 4
                arr2[i][j] = min(max(round(arr2[i][j] + 128), 0), 255)


        print_arr(arr2)
