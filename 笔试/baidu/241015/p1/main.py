# 从1～n中选k个，得分 = 所有(i+1)没选的i数量，问最大得分

T = int(input())

for _ in range(T):
    n, k = map(int, input().split())
    if n == k:
        print(1)
    # elif n==k+1:
    #     print(1)
    else:
        ans = 1
        n -= 1
        k -= 1

        pair = n // 2
        if k <= pair:
            ans += k
        else:
            ans += pair
            sub_val = k - pair
            if n%2==1:
                sub_val -= 1
            ans -= sub_val

        print(ans)


# 0 1 0 1 0 1