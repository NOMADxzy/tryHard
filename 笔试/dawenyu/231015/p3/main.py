from scipy.sparse.linalg import spilu

n, m = map(int, input().split())
labels = list(map(int, input().split()))
# 5 6
# 1 2 1 2 3
# 1 1 5
# 3 1
# 1 2 7
# 3 4
# 2 4 4
# 3 2

groups = {}
for label in labels:
    if label in groups:
        groups[label][1] += 1
    else:
        groups[label] = [0, 1]



for _ in range(m):
    splits = input().split()
    if len(splits) == 3:
        opt, idx, val = map(int, splits)
    else:
        opt, idx = map(int, splits)
    idx -= 1
    label = labels[idx]
    if opt == 1:
        groups[label][0] +=  val
    elif opt == 2:
        groups[label][0] -= val
    else:
        total, size = groups[label]
        print(total / size)


