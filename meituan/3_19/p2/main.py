import sys

n,q = 0,0

def solve(sum,k,l,r):
    return sum + k*l, sum + k*r

t = 0
sum,k = 0,0
for line in sys.stdin:
    if t==0:
        a = line.split()
        n = int(a[0])
        q = int(a[1])
    elif t==1:
        a = line.split()
        sum = 0
        k = 0
        for i in range(0, len(a)):
            if int(a[i]) == 0:
                k += 1
            else:
                sum += int(a[i])
    else:
        a = line.split()
        l, r = int(a[0]), int(a[1])
        minVal, maxVal = solve(sum, k, l, r)
        print(minVal, maxVal)
    t += 1
    if t==q+2:
        break

