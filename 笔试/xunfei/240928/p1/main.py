T = int(input())


def count1(idx: int) -> int:
    ans = 0
    while idx > 0:
        ans += 1
        idx -= idx & (-idx)

    return ans


for _ in range(T):
    s = input()
    new_s = ""
    for i in range(len(s)):
        cur = s[i]
        if count1(i + 1) % 2 == 1:
            cur = chr(ord(cur) - 32)
        new_s += cur

    print(new_s)
