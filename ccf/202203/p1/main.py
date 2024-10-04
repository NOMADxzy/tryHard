
n, k = map(int, input().split())

ans = 0
initialized = set()

for _ in range(k):
    x, y = map(int, input().split())
    if y>0 and y not in initialized:
        ans += 1
    initialized.add(x)

print(ans)