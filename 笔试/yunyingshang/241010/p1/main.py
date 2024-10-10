
T = int(input())

for _ in range(T):
    h, a, H, A = map(int, input().split())
    # 5 1 2 1
    ans = 0
    attacked_time = H // a - 1
    if H % a > 0:
        attacked_time += 1
    if attacked_time * A < h:
        health_sub_per_kill = attacked_time * A
        if health_sub_per_kill == 0:
            ans = -1
        else:
            ans = h // health_sub_per_kill
            if h % health_sub_per_kill == 0:
                ans -= 1
    print(ans)

