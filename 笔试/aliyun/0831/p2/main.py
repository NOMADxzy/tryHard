n, l, s = map(int, input().split())

# 5 5 13

# 0101
# 0110
# 0111
# 1000
# 1001
#
#
# 0101
# 1101

origin_xor_val = 0
for x in range(l, l+n):
    origin_xor_val ^= x

if origin_xor_val == s:
    print(0)
else:
    tmp_xor_val = origin_xor_val ^ l
    new_l_val = tmp_xor_val ^ s
    print(l, new_l_val)

