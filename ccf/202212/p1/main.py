
n, q = input().split()
n = int(n)
q = float(q)
rewards = list(map(int, input().split()))

w = 1 + q
cur_q = w
acc = rewards[0] * 1
for i in range(1, len(rewards)):
    acc += rewards[i] / cur_q
    cur_q *= w

print(acc)