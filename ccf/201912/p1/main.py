n = int(input())


def has_pre_7(x: int) -> bool:
    mask = 10
    while mask <= x:
        i = (x // mask) % 10
        if i == 7:
            return True
        mask *= 10
    return False


ans = [0, 0, 0, 0]
num_idx = 0
# pos_idx = 0
i = 1
pre_has_7 = False
while num_idx < n:
    jump = False
    if i % 7 == 0 or i % 10 == 7:
        jump = True
        if i % 10 == 0:
            pre_has_7 = has_pre_7(i)
    else:
        if i % 10 == 0:
            pre_has_7 = has_pre_7(i)
        if pre_has_7:
            jump = True
    if jump:
        ans[(i - 1) % 4] += 1
    else:
        num_idx += 1
    i += 1
# print(i-1)

print("\n".join(map(str, ans)))
