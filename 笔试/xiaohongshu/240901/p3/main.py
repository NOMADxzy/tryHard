
# 10
# RRBBBBBBBBBB
# 1 2
# 1 3
# 1 4
# 1 5
# 1 7
# 2 8
# 4 6
# 1 9
# 1 10

n = int(input())
s = input()
totalR = 0
for c in s:
    if c=='R':
        totalR += 1

nexts = [[] for _ in range(n+1)]
for i in range(n-1):
    a, b = map(int, input().split())
    nexts[a].append(b)
    nexts[b].append(a)


def count_b_num(pos: int, last_pos: int) -> int:
    cnt = 1 if s[pos] == 'B' else 0
    for np in nexts[pos]:
        if np != last_pos:
            cnt += count_b_num(np, pos)

    return cnt

ans = 0
for i in range(n):
    if s[i] == 'R':
        idx = i+1
        for np in nexts[idx]:
            ans = max(ans, count_b_num(pos=np, last_pos=idx))

print(ans)
