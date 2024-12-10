# 5 2 6 20 40 100 80
# 100 80 100 80 100 80 100 80 100 80 100 80
# 60 50 60 46 60 42 60 38 60 34 60 30
# 10 60 14 62 18 66 22 74 26 86 30 100
# 90 31 94 35 98 39 102 43 106 47 110 51
# 0 20 4 20 8 20 12 20 16 20 20 20

n, k, t, x1, y1, x2, y2 = map(int, input().split())
cnt_inside, cnt_inside_long = 0, 0


def inside(x: int, y: int) -> bool:
    return x1 <= x <= x2 and y1 <= y <= y2


for _ in range(n):
    input_ints = list(map(int, input().split()))
    acc_inside = 0
    is_inside = False
    is_inside_long = False
    for i in range(0, 2 * t, 2):
        cur_x, cur_y = input_ints[i], input_ints[i + 1]
        if inside(cur_x, cur_y):
            acc_inside += 1
            is_inside = True
        else:
            acc_inside = 0
        if acc_inside >= k:
            is_inside_long = True
            break
    if is_inside:
        cnt_inside += 1
    if is_inside_long:
        cnt_inside_long += 1

print(cnt_inside, cnt_inside_long)
