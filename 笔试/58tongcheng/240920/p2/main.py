
s = input()

cnta, cntb = 0, 0
for c in s:
    if c == 'b':
        cntb += 1

if s[0] == 'a':
    cnta += 1
else:
    cntb -= 1

n = len(s)
ans = cnta + cntb
for i in range(1, n-1):
    if s[i] == 'a':
        cnta += 1
    else:
        cntb -= 1
    ans = max(ans, cnta + cntb)

print(ans)