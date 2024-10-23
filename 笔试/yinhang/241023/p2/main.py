# 10 10 5
# 1 1 3 10
# 1 1 3 2
# 1 2 4 3
# 2 1 4 5
# 1 3 5 1

# 17
# k个地鼠，每个地鼠会在t时刻从(x,y)位置钻出来，得分s，初始时在左上角，每个单位时间可移动一格，问最多能得多少分

n, m, k = map(int, input().split())

for _ in range(k):
    x, y, t, s = map(int, input().split())