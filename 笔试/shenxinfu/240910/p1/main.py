n, m = map(int, input().split())

for _ in range(m):
    s = input()
    ans = 0

    i = 0
    cnts = {}
    flag = False
    while i < len(s):
        j = i
        while j < len(s) and s[j] == s[i]:
            j += 1
        if s[i] not in cnts:
            cnts[s[i]] = 0
        else:
            flag = True
        cnts[s[i]] += 1
        i = j

    if not flag:
        ans = 1

    print(ans)