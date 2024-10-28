from typing import List
from math import pow

T = int(input())
# 5
# 10
# 15
# 123
# 33
# 321

def dfs(num: int) -> List[List[int]]:
    if num == 0:
        return []
    exp2, exp3 = 0, 0
    while num % 2 == 0:
        exp2 += 1
        num //= 2
    while num % 3 == 0:
        exp3 += 1
        num //= 3
    ret = [[exp2, exp3]]
    next_ret = dfs(num-1)
    for e in next_ret:
        ret.append([e[0]+exp2, e[1]+exp3])
    return ret

for _ in range(T):
    num = int(input())
    pairs = dfs(num)
    ans = [2**p[0] * 3**p[1]  for p in pairs]
    ans.sort(reverse=True)
    print(*ans)
