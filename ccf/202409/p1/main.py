
# 4
# csp#ccsp
# csp#ccsp2024
# Csp#ccsp2024
# CSP#2024

n = int(input())

def judge(s: str) -> int:
    cnts = {}
    has_en, has_num, has_spec = False, False, False
    more_than_2 = False
    for c in s:
        if c not in cnts:
            cnts[c] = 0
        cnts[c] += 1
        if cnts[c]>2:
            more_than_2 = True
        if ord(c)>=65 and ord(c)<=90 or ord(c)>=97 and ord(c)<=122:
            has_en = True
        elif ord(c)>=48 and ord(c)<=57:
            has_num = True
        else:
            has_spec = True
    if not (has_en and has_num and has_spec):
        return 0
    elif more_than_2:
        return 1
    else:
        return 2

for _ in range(n):
    print(judge(input()))