from math import inf
T = int(input())

# 1 3 5 2 4
# 5 2 3 4 1

for _ in range(T):
    n = int(input())
    a_arr = list(map(int, input().split()))
    b_arr = list(map(int, input().split()))

    dp_up = [[True, True] for _ in range(n)]
    dp_down = [[True, True] for _ in range(n)]

    for i in range(1, n):
        best_val_0 = False
        if a_arr[i] >= a_arr[i-1] and dp_up[i-1][0]:
            best_val_0 = True
        if a_arr[i] >= b_arr[i-1] and dp_up[i-1][1]:
            best_val_0 = True
        dp_up[i][0] = best_val_0

        best_val_1 = False
        if b_arr[i] >= b_arr[i - 1] and dp_up[i-1][1]:
            best_val_1 = True
        if b_arr[i] >= a_arr[i - 1] and dp_up[i-1][0]:
            best_val_1 = True
        dp_up[i][1] = best_val_1

    for i in range(1, n):
        best_val_0 = False
        if a_arr[i] <= a_arr[i-1] and dp_down[i-1][0]:
            best_val_0 = True
        if a_arr[i] <= b_arr[i-1] and dp_down[i-1][1]:
            best_val_0 = True
        dp_down[i][0] = best_val_0

        best_val_1 = False
        if b_arr[i] <= b_arr[i - 1] and dp_down[i-1][1]:
            best_val_1 = True
        if b_arr[i] <= a_arr[i - 1] and dp_down[i-1][0]:
            best_val_1 = True
        dp_down[i][1] = best_val_1

    if dp_up[n-1][0] or dp_up[n-1][1] or dp_down[n-1][0] or dp_down[n-1][1]:
        print("YES")
    else:
        print("NO")

