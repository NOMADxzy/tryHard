# 求1到l位n进制 所能形成的id数目
MOD = 1000000007

def cpt(n, l):
    if l == 0:
        return 1
    elif l == 1:
        return n
    else:
        t = 1
        if l%2==1:
            t *= n
        x = cpt(n, l//2)
        ret = (((x * x) % MOD) * t) % MOD
        return ret

def mod_exp_sum(n, l):
    MOD = 1000000007
    result = 0
    current_power = 1

    for i in range(1, l + 1):
        current_power = (current_power * n) % MOD
        result = (result + current_power) % MOD

    return result

while True:
    n, l = map(int, input().split())
    if n==0 and l==0:
        break


    # 输出结果
    print(mod_exp_sum(n, l))


