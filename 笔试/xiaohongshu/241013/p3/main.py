# 能否最多交换一次的情况下，使字符串变成周期串

T = int(input())

for _ in range(T):
    s = input()
    cnts = [0] * 26
    for c in s:
        cnts[ord(c) - 97] += 1

    t = max(cnts)
    ok = True
    for x in cnts:
        if x > 0 and x != t:
            ok = False
            break

    print("YES" if ok else "NO")