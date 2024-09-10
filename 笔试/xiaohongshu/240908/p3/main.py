from numpy.ma.core import append

n = int(input())
arr = list(map(int, input().split()))

exsits = {}
new_arr = []
for v in arr:
    if v not in exsits:
        exsits[v] = True
        new_arr.append(v)

arr = new_arr
arr.sort()
n = len(arr)




def g(x):
    mask = 1
    while mask<=x and x&mask==0:
        mask *= 2

    while mask<=x and (x&mask)>0:
        mask *= 2

    return x + mask//2

dp = {}
def dfs(x):
    if x in dp:
        return dp[x]
    if x not in exsits:
        return 0
    tmp = g(x)
    ret = 1 + dfs(tmp)
    dp[x] = ret
    return ret

ans = 0
for val in arr:
    ans = max(ans, dfs(val))

print(ans)
# print(g(3))

