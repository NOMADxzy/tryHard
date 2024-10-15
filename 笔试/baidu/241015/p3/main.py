
# 交替+-，
# 原始: 1 2 3 4
# 1+2, 2-3, 3+4  ->  3, -1, 7
# 3-(-1), -1+7   ->  4 6
# 4-6            ->  -2

n = int(input())
arr = list(map(int, input().split()))

MOD = 1000000007

first_symb = 1
while len(arr)>1:
    new_arr = []
    symb = first_symb
    for i in range(len(arr)-1):
        x = (arr[i] + symb * arr[i+1]) % MOD
        new_arr.append(x)
        symb = -symb
    first_symb = symb
    arr = new_arr

print(arr[0])