# 1到100000000个价值的礼物，已经有其中的input_arry，还有money的钱，问还能买多少个礼物
def solution(input_arry, money, giftNum):
    # 在此函数内完成题目要求，可定义其他函数
    pre_sums = [0] * (giftNum + 1)
    input_arry.sort()
    for i in range(giftNum):
        pre_sums[i + 1] = pre_sums[i] + input_arry[i]
    left, right = 1, 1000000000

    def check(cur_idx):
        total_money = (1 + cur_idx) * cur_idx // 2
        already_money = 0
        already_num = 0
        l, r = 0, giftNum - 1
        if input_arry[r] <= cur_idx:
            already_money = pre_sums[-1]
            already_num = r + 1
        elif input_arry[l] > cur_idx:
            pass
        else:
            while l < r:
                mid = (l + r + 1) // 2
                if input_arry[mid] <= cur_idx:
                    l = mid
                else:
                    r = mid - 1
            already_money = pre_sums[l + 1]
            already_num = l + 1
        return total_money - already_money <= money, cur_idx - already_num

    if not check(left):
        return 0
    while left < right:
        mid = (left + right + 1) // 2
        ok, _ = check(mid)
        if ok:
            left = mid
        else:
            right = mid - 1
    _, ans = check(left)
    return ans


# 1 3 4 6 7

if __name__ == "__main__":
    n, m = map(int, input().split());
    if (n == 0):
        input_arry = {0}
    else:
        input_arry = list(map(int, input().split(' ')));
    print(solution(input_arry, m, n))