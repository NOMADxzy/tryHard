from math import pow
from copy import deepcopy

w, s = map(int, input().split())
origin = input()

data = []
mode = 0 # 大写、小写、数字

def encode(x: str) -> (int, int):
    asc_val = ord(x)
    if asc_val >= 97 and asc_val <= 122:
        return asc_val - 97, 1
    elif asc_val >= 65 and asc_val <= 90:
        return asc_val - 65, 0
    else:
        return asc_val - ord('0'), 2

def add_change(pre_mode, new_mode):
    if pre_mode == new_mode:
        return
    elif pre_mode == 0:
        if new_mode == 1:
            data.append(27)
        else:
            data.append(28)
    elif pre_mode == 1:
        if new_mode == 2:
            data.append(28)
        else:
            data.append(28)
            data.append(28)
    else:
        if new_mode == 0:
            data.append(28)
        else:
            data.append(27)

for c in origin:
    encoded_val, new_mode = encode(c)
    add_change(pre_mode=mode, new_mode=new_mode)
    data.append(encoded_val)
    mode = new_mode

if len(data) % 2 == 1:
    data.append(29)

k = 1<<(s+1) if s>=0 else 0
data_nums, pad_nums, crc_nums = [0], [], [0 for _ in range(k)]
for i in range(0, len(data), 2):
    data_nums.append(30 * data[i] + data[i+1])

pad_len = w - (len(data_nums) + len(crc_nums)) % w
for _ in range(pad_len):
    pad_nums.append(900)

data_nums[0] = len(data_nums) + len(pad_nums)

dx = data_nums + pad_nums
dx.reverse() # 0~n-1 -> k~k+n-1
n = len(dx)

def mul(list1, list2):
    m, n = len(list1), len(list2)
    cnts = [0 for _ in range((m-1) + (n-1) + 1)]
    for i in range(m):
        for j in range(n):
            new_idx = i + j
            cnts[new_idx] += list1[i] * list2[j]
    return cnts

def add(list1, list2):
    if len(list1) < len(list2):
        list1, list2 = list2, list1
    m, n = len(list1), len(list2)

    for i in range(n):
        if i<k:
            list1[i] = (list1[i] + list2[i] % 929) % 929
        else:
            list1[i] = list1[i] + list2[i]


    return list1

if k>0:
    gx = [1] # 0~k
    for i in range(1, k+1):
        gx = mul(gx, [-int(pow(3, i)), 1])

    qx = [0 for _ in range(n)]
    qgx = [0 for _ in range(k+n)]
    j = k
    tmp_qx = [0 for _ in range(n)]
    for i in range(n-1,-1,-1):
        qx[i] = dx[i] - qgx[i+k]
        tmp_qx[i] = qx[i]

        qgx = add(deepcopy(qgx), mul(tmp_qx, gx))
        tmp_qx[i] = 0

    rx = qgx[:k]

    for i in range(0, k):
        crc_nums[i] = rx[k-1-i] % 929

for x in data_nums:
    print(x)

for x in pad_nums:
    print(x)

for x in crc_nums:
    print(x)


