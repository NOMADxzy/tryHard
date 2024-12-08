from typing import Dict


# 11
# H2+O2=H2O
# 2H2+O2=2H2O
# H2+Cl2=2NaCl
# H2+Cl2=2HCl
# CH4+2O2=CO2+2H2O
# CaCl2+2AgNO3=Ca(NO3)2+2AgCl
# 3Ba(OH)2+2H3PO4=6H2O+Ba3(PO4)2
# 3Ba(OH)2+2H3PO4=Ba3(PO4)2+6H2O
# 4Zn+10HNO3=4Zn(NO3)2+NH4NO3+3H2O
# 4Au+8NaCN+2H2O+O2=4Na(Au(CN)2)+4NaOH
# Cu+As=Cs+Au

def is_upper(c):
    return 65 <= ord(c) <= 90


def is_lower(c):
    return 97 <= ord(c) <= 122


def is_num(c):
    return 48 <= ord(c) <= 57


def get_all_cnts(s: str) -> Dict:
    total_times = 1
    i = 0
    if ord(s[i]) < 58:
        while ord(s[i]) < 58:
            i += 1
        total_times = int(s[:i])
    cnts = {}
    while i < len(s):
        if is_upper(s[i]):
            j = i + 1
            if j < len(s) and is_lower(s[j]):
                j += 1
            cur_elem = s[i:j]
            cur_times = 1
            i = j
            while j < len(s) and is_num(s[j]):
                j += 1
            if j > i:
                cur_times = int(s[i:j])
            i = j
            if cur_elem not in cnts:
                cnts[cur_elem] = 0
            cnts[cur_elem] += cur_times
        elif s[i] == '(':
            left_cnt = 1
            j = i + 1
            while left_cnt > 0:
                if s[j] == ')':
                    left_cnt -= 1
                elif s[j] == '(':
                    left_cnt += 1
                j += 1
            inner_cnts = get_all_cnts(s[i + 1:j - 1])
            cur_times = 1
            i = j
            while j < len(s) and is_num(s[j]):
                j += 1
            if j > i:
                cur_times = int(s[i:j])
            i = j
            for inner_elem, inner_cnt in inner_cnts.items():
                if inner_elem not in cnts:
                    cnts[inner_elem] = 0
                cnts[inner_elem] += inner_cnt * cur_times
        else:
            raise ValueError(s[i])
    for k, v in cnts.items():
        cnts[k] = total_times * v
    return cnts


n = int(input())
for _ in range(n):
    left, right = input().split("=")
    left_elems = left.split("+")
    right_elems = right.split("+")
    left_total_cnts, right_total_cnts = {}, {}
    for left_elem in left_elems:
        tmp = get_all_cnts(left_elem)
        for k, v in tmp.items():
            left_total_cnts[k] = left_total_cnts[k] + v if k in left_total_cnts else v
    for right_elem in right_elems:
        tmp = get_all_cnts(right_elem)
        for k, v in tmp.items():
            right_total_cnts[k] = right_total_cnts[k] + v if k in right_total_cnts else v

    is_ok = len(left_total_cnts) == len(right_total_cnts)
    if is_ok:
        for k, v in left_total_cnts.items():
            if k not in right_total_cnts or right_total_cnts[k] != v:
                is_ok = False

    print("Y" if is_ok else "N")
