
n = int(input())
s = input()

# 01100100

enc_str = "0X"
enc_vals = []

b_arr = []
mask = 1
while mask <= n:
    if mask & n > 0:
        b_arr.append(1)
    else:
        b_arr.append(0)
    mask *= 2

dict_0x = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F']
def get_val(s):
    if ord(s) >= 65:
        return ord(s) - 65 + 10
    else:
        return int(s)

def get_arr(s):
    t1, t2 = get_val(s[0]), get_val(s[1])
    t = t1*16 + t2
    arr = [0 for _ in range(7)]
    mask = 1
    for i in range(7):
        if t & mask>0:
            arr[i] = 1
        mask *= 2
    return arr

def trans_to_0x(arr):
    mask = 1
    acc = 0
    for x in arr:
        acc += mask * x
        mask *= 2
    return dict_0x[acc]
# print(b_arr)

for i in range(0, len(b_arr), 7):
    is_last = i+7>=len(b_arr)
    cur1, cur2 = b_arr[i:i+4], b_arr[i+4:i+7]
    cur2.append(0 if is_last else 1)
    # print(cur1, cur2)
    enc_vals.append(trans_to_0x(cur2) + trans_to_0x(cur1))
# print(enc_vals)
enc_vals.reverse()
enc_str += ''.join(enc_vals)


dec_b_arr = []
for i in range(0, len(s), 2):
    if s[i:i+2] == '0X':
        continue
    tmp_arr = get_arr(s[i:i+2])
    # print(tmp_arr)
    dec_b_arr.extend(tmp_arr)

# print(dec_b_arr)
dec_val = 0
mask = 1
for x in dec_b_arr:
    dec_val += mask * x
    mask *= 2

print(enc_str)
print(dec_val)




