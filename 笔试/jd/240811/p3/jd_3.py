
# 3 6 2
# 1 2 1
# 1 2 1
# 1 2 1
# 2 3 1
# 2 3 1
# 2 3 1
#

n, line_cnt, target = map(int, input().split())

grid = [[{} for _ in range(n)] for _ in range(n)]


pres = [{} for i in range(n)]
for _ in range(line_cnt):
    f, t, cost = map(int, input().split())
    f-=1
    t-=1
    m = grid[f][t]
    if cost in m:
        m[cost] += 1
    else:
        m[cost] = 1
    pres[t][f] = True


MOD = 20220201
moded = False
def dfs(m, cost)->int:
    if m==0:
        return 1 if cost==0 else 0
    cnt = 0
    for p in pres[m].keys():
        for c, times in grid[p][m].items():
            if cost - c >= 0:
                v = times * dfs(p, cost-c)
                if cnt + v > MOD:
                    moded = True
                cnt = (cnt + v) % MOD

    return cnt


ans = dfs(n-1, target)
if moded:
    print("All roads lead to Home")

print(ans)