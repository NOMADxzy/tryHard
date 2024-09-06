

n, k = map(int, input().split())

ans = []

for i in range(1, k):
    ans.append(i)

ans.append(n)

for i in range(n-1, k-1, -1):
    ans.append(i)

print(" ".join(map(str, ans)))
