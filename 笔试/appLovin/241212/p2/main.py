# 将字符串中所有字符ascii前移一位
# abc XYZ

s = input()


def preChar(c: str) -> str:
    asc_idx = ord(c)
    if 65 <= asc_idx <= 90:
        new_idx = asc_idx - 1
        if new_idx < 65:
            new_idx += 26
        return chr(new_idx)
    elif 97 <= asc_idx <= 122:
        new_idx = asc_idx - 1
        if new_idx < 97:
            new_idx += 26
        return chr(new_idx)
    else:
        return c


ans = ""
for c in s:
    ans += preChar(c)

print(ans)
