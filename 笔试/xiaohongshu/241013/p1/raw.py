from math import ceil
m, n = map(int, input().split())


def get_score(c):
    if ord(c) <= 90:
        return ord(c) - 65
    else:
        return ord(c) - 97

c_dict = [chr(i + 97) for i in range(26)]

score_sums = [0] * n
score_max = [0] * n
score_min = [25] * n
for _ in range(m):
    s = input()
    for i in range(n):
        val = get_score(s[i])
        score_max[i] = max(score_max[i], val)
        score_min[i] = min(score_min[i], val)
        score_sums[i] += val

pairs = []
for i in range(n):
    mean_val = ceil((score_sums[i] - score_max[i] - score_min[i]) // (m-2))
    pairs.append((c_dict[mean_val], i+1))

pairs.sort()
row1 = [e[1] for e in pairs]
row2 = [e[0] for e in pairs]

print(" ".join(map(str, row1)))
print(" ".join(row2))

