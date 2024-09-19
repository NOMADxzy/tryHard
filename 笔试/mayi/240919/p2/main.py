
# 给一个字符串，删除多少个字符可以让其称为非回文字符串，或者单个字符
# 输出最小的删除次数

T = int(input())

for _ in range(T):
    n = int(input())
    s = input()
    ok = True
    c_map = {}
    for i in range(n // 2):
        if s[i] != s[n - 1 - i]:
            ok = False
            break
        else:
            if s[i] not in c_map:
                c_map[s[i]] = 0
            c_map[s[i]] += 1
    if not ok:
        print(0)
    elif len(c_map) > 1:
        print(1)
    else:
        print(n - 1)


