
n, m, k = map(int, input().split())

plans = []
max_t = 0
for _ in range(n):
    tmp = list(map(int, input().split()))
    max_t = max(max_t, tmp[0])
    plans.append(tmp)

activities = [0 for _ in range(max_t+2)]
for t, c in plans:
    activities[max(0, t-c+1)] += 1
    activities[t+1] -= 1

acc = 0
for i in range(len(activities)):
    activities[i] += acc
    acc = activities[i]

for _ in range(m):
    q = int(input())
    if q+k <= max_t:
        print(activities[q+k])
    else:
        print(0)


