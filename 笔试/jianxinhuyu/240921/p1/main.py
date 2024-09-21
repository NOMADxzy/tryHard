
# 每一行、每一列、每条对角线形成的连续1的长度最大为多少？
# 5 5
# 0 0 0 0 1
# 1 1 0 1 0
# 0 1 1 0 0
# 0 1 0 1 0
# 1 0 0 0 1
arr = []
m,n = map(int, input().split())
for _ in range(m):
    arr.append(list(map(int, input().split())))

rows = [0 for _ in range(m)]
cols = [0 for _ in range(n)]
ex1s = [0 for _ in range(m+n-1)]
ex2s = [0 for _ in range(m+n-1)]

for i in range(m):
    for j in range(n):
        if arr[i][j] == 1:
            rows[i] += 1
            cols[j] += 1
            ex1_idx = i+j
            ex2_idx = m-1-i+j
            ex1s[ex1_idx] += 1
            ex2s[ex2_idx] += 1

ans = max(max(rows), max(cols), max(ex1s), max(ex2s))
print(ans)
