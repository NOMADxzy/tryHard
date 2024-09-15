
n, a, b = map(int, input().split())

# 4 10 10
# 0 0 5 5
# 5 -2 15 3
# 8 8 15 15
# -2 10 3 15

ans = 0
for _ in range(n):
    x1, y1, x2, y2 = map(int, input().split())
    if x2 >0 and y2 > 0 and x1 < a and y1 < b:
        t1 = max(x1, 0)
        s1 = max(y1, 0)
        t2 = min(x2, a)
        s2 = min(y2, b)
        ans += (t2 - t1) * (s2 - s1)

print(ans)
