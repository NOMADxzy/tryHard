
n, m = map(int, input().split())
prices = list(map(int, input().split()))

# 9 6
# 2 3 3 6 6 6 9 9 23

max_key, max_val = -1, 0
cnts = {}
for p in prices:
    if p not in cnts:
        cnts[p] = 1
    else:
        cnts[p] += 1
    if cnts[p] > max_val:
        max_key = p
        max_val = cnts[p]

print(max_val)