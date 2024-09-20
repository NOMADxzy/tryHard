
n, m = map(int, input().split())

sum_x, sum_y = 0, 0

for _ in range(0, n):
    dx, dy = map(int, input().split())
    sum_x += dx
    sum_y += dy

for _ in range(m):
    x, y = map(int, input().split())
    print(sum_x + x, sum_y + y)