from typing import List, Tuple
from math import inf
# TODO
T = int(input())
# T = 1
# 5 5
# 1 5
# 1 2 1
# 1 3 6
# 2 3 2
# 2 5 5
# 2 4 10

class Path:
    def __init__(self):
        self.paths = []

def better(a1, b1, a2, b2):
    return a1*b2 < a2*b1

paths = []
def get_best(pos: int, target: int, nexts: List[List[Tuple]], state_str: List[str], dp: {}, paths: Path)->(int, int):
    max_val, min_val = inf, 1
    state_str[pos] = '1'
    paths.paths.append(pos)
    key = "_".join(map(str, paths.paths))
    if key in dp:
        return dp[key]

    for np, w in nexts[pos]:
        if state_str[np]=="1":
            continue
        if np == target :
            max_val, min_val = w, w
        else:
            next_max_val, next_min_val = get_best(np, target, nexts, state_str, dp, paths)
            next_max_val = max(w, next_max_val)
            next_min_val = min(w, next_min_val)
            if better(next_max_val, next_min_val, max_val, min_val):
                max_val, min_val = next_max_val, next_min_val
    dp[key] = (max_val, min_val)
    state_str[pos] = '0'
    paths.paths = paths.paths[:-1]
    return max_val, min_val


for _ in range(T):
    n, m = map(int, input().split())
    nexts = [[] for _ in range(n+1)]
    s, t = map(int, input().split())
    for _ in range(m):
        x, y, w = map(int, input().split())
        nexts[x].append((y, w))
        nexts[y].append((x, w))

    max_val, min_val = get_best(s, t, nexts, ["0" for _ in range(n+1)], {}, Path())
    if max_val != inf:
        out = str(max_val//min_val) if max_val%min_val == 0 else f'{max_val}/{min_val}'
        print(out)
    else:
        print("%%%")