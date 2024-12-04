
# 2 3 3 2
# 1 2 3
# 4 5 6

m, n, p, q = map(int, input().split())
flatten_arr = []
for _ in range(m):
    flatten_arr.extend(map(int, input().split()))

ans = []
for i in range(0, p*q, q):
    ans.append(flatten_arr[i:i+q])

for row in ans:
    print(" ".join(map(str, row)))
