
# 1.2.3
# 1.2.4

v1 = list(map(int, input().split(".")))
v2 = list(map(int, input().split(".")))

n = len(v1)
ans = 0
for i in range(n):
    if v1[i]>v2[i]:
        ans = 1
        break
    elif v1[i]<v2[i]:
        ans = 2
        break
    else:
        continue

print(ans)