from math import inf

n = int(input())
book_types = list(map(int, input().split()))
can_move = list(map(lambda x: x == '1', input().split()))

remain1, remain2 = 0, 0
for i, e in enumerate(book_types):
    if not can_move[i]:
        continue
    if e == 1:
        remain1 += 1
    else:
        remain2 += 1

total = remain1 + remain2

# 5
# 1 2 1 2 1
# 0 1 1 0 1
# 1 . . 2 .

dp = {}


def dfs(pos: int, last_book_type: int, remain1: int, remain2: int) -> int:
    if pos == n:
        return 0
    key = remain1 * (total + 1) + remain2
    if last_book_type == 2:
        key = -key
    if key in dp:
        return dp[key]
    min_val = inf
    if can_move[pos]:
        assert remain1 > 0 or remain2 > 0
        if remain1 > 0:
            add_val = 1 if last_book_type == 2 else 0
            min_val = min(min_val, dfs(pos + 1, 1, remain1 - 1, remain2) + add_val)
        if remain2 > 0:
            add_val = 1 if last_book_type == 1 else 0
            min_val = min(min_val, dfs(pos + 1, 2, remain1, remain2 - 1) + add_val)
    else:
        add_val = 1 if pos > 0 and book_types[pos] != last_book_type else 0
        min_val = min(min_val, dfs(pos + 1, book_types[pos], remain1, remain2) + add_val)

    dp[key] = min_val
    return int(min_val)


print(dfs(0, 0, remain1, remain2))
