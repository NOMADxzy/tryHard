articles = list(map(int, input().split()))
users = list(map(int, input().split()))
boosts = int(input())
b_value = int(input())
max_use = int(input())

users.sort(reverse=True)
articles.sort(reverse=True)
n = len(articles)

left, right = 0, n


def check(x):
    if x > len(users):
        return False

    arts = articles[n - x:]
    use_b = 0
    i = 0
    while use_b <= boosts and i < x:
        dif = arts[i]
        uds = users[i]
        t = 0
        while uds < dif and t <= max_use and use_b <= boosts:
            t += 1
            uds += b_value
            use_b += 1
        if t > max_use or use_b > boosts:
            return False
        i += 1
    return i == x


if check(right):
    print(right)
else:

    while left < right:
        mid = (left + right + 1) // 2
        if check(mid):
            left = mid
        else:
            right = mid - 1
    print(left)
