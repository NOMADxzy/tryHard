# 给一个01字符串，进行n次字符位置交换，得到的最小字典序的字符串是什么

T = int(input())

for _ in range(T):
    n, k = map(int, input().split())
    s = input()
    left, right = 0, n-1
    ans = [c for c in s]
    while right>left and s[right]=='1':
        right -= 1

    if s[right] == '1':
        print(s)
    elif s == '01' or s == '10':
        if s == '01' and k%2 == 0 or s == '10' and k%2 == 1:
            print('01')
        else:
            print('10')
    else:
        while left < right and k>0:
            while left < right and s[left] == '0':
                left += 1
            if left == right:
                break
            ans[right] = '1'
            ans[left] = '0'
            k -= 1
            left += 1
            right -= 1
            while right > left and s[right] == '1':
                right -= 1

        print("".join(ans))
