
n = int(input())
has_rabbish = {}
poses = []
for _ in range(n):
    x, y = map(int, input().split())
    key_str = f"{x}_{y}"
    has_rabbish[key_str] = True
    poses.append([x, y])

dirs1 = [[-1,0],[1,0],[0,-1],[0,1]]
dirs2 = [[-1,-1],[1,-1],[-1,1],[1,1]]
ans = [0]*5
for x,y in poses:
    ok = True
    for dx, dy in dirs1:
        key_str = f"{x+dx}_{y+dy}"
        if key_str not in has_rabbish:
            ok = False
            break
    if ok:
        cnt = 0
        for dx, dy in dirs2:
            key_str = f"{x + dx}_{y + dy}"
            if key_str in has_rabbish:
                cnt += 1
        ans[cnt] += 1

print("\n".join(map(str, ans)))

# 11
# 9 10
# 10 10
# 11 10
# 12 10
# 13 10
# 11 9
# 11 8
# 12 9
# 10 9
# 10 11
# 12 11

# 0
# 2
# 1
# 0
# 0
