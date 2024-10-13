# 每行表示评委对每个选手的评分，选手得分=去掉最大最小值的均值 ，求最终选手得分从大到小的编号、分值

from math import ceil
n, m = map(int, input().split())


def get_score(c):
    if ord(c) <= 90:
        return ord(c) - 65
    else:
        return ord(c) - 97


score_of_candidates = [[] for i in range(m)]
for _ in range(n):
    s = input()
    for i in range(m):
        score_of_candidates[i].append(get_score(s[i]))

scores = []

c_dict = [chr(i + 97) for i in range(26)]

for idx in range(m):
    raw_score = score_of_candidates[idx]
    min_val, max_val = min(raw_score), max(raw_score)
    mean_val = ceil((sum(raw_score) - min_val - max_val) / (n - 2))
    # mean_scores.append(c_dict[mean_val])
    scores.append((mean_val, idx + 1))

scores.sort()
idxs = [e[1] for e in scores]
mean_scores = [c_dict[e[0]] for e in scores]


print(" ".join(map(str, idxs)))
print(" ".join(map(str, mean_scores)))
