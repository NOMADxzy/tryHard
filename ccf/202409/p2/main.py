
# #Hello World#
# 6
# #HH#
# #e #
# # r#
# #re#
# #oa#
# #ao#
# 3
# 1 2 3

origin = input()[1:-1]
f_cnt = int(input())
f_dict = {}

circles = []
circle_pos = {}

for _ in range(f_cnt):
    tmp = input()[1:-1]
    f, t = tmp[0], tmp[1]
    f_dict[f] = t

def f_func(x):
    if x not in f_dict:
        return x
    else:
        return f_dict[x]

for f in origin:
    visited = {f:0}
    t = f_func(f)
    path = f
    while t not in circle_pos:
        if t in visited:
            start_pos, end_pos = visited[t], visited[f]
            circles.append(path[start_pos: end_pos+1])
            for i in range(start_pos, end_pos+1):
                circle_pos[path[i]] = [len(circles)-1, i-start_pos]
            break
        else:
            visited[t] = len(path)
            path += t
            f, t = t, f_func(t)

m = int(input())
k_list = list(map(int, input().split()))

def cpt(x, k):
    while x not in circle_pos and k>0:
        x = f_func(x)
        k -= 1
    if k==0:
        return x
    else:
        circle_idx, internal_idx = circle_pos[x]
        path = circles[circle_idx]
        return path[(internal_idx+k)%len(path)]

for k in k_list:
    ans = ""
    for c in origin:
        ans += cpt(c, k)
    print(ans)