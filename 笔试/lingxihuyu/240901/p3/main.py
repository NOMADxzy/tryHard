
n = int(input())
weights = list(map(int, input().split()))

# 5
# 1 2 3 17 10

ans = 0
weights.sort()

i = 0
cur = weights[0]
while i<n:
    while i<n and weights[i] <= cur + 4:
        i += 1

    ans += 1
    cur = weights[i] if i<n else cur

print(ans)
