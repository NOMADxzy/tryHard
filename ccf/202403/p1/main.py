# 4 3
# 5 1 2 3 2 1
# 1 1
# 3 2 2 2
# 2 3 2

n, m = map(int, input().split())
cnts = [0]*m
exsits = [set() for _ in range(m)]
for article_idx in range(n):
    arr = list(map(int, input().split()))
    arr = arr[1:]
    for w_idx in arr:
        w_idx -= 1
        cnts[w_idx] += 1
        exsits[w_idx].add(article_idx)

for w_idx in range(m):
    print(len(exsits[w_idx]), cnts[w_idx])