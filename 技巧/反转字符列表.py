# ['a', 'b', 'c', ' ', ' ', 'd', 'e', ' ', 'f'] ->  ['f', ' ', 'e', 'd', ' ', ' ', 'c', 'b', 'a', ]

def reverse(arr: list) -> list:
    n = len(arr)
    for i in range(n // 2):
        arr[i], arr[n - 1 - i] = arr[n - 1 - i], arr[i]

    i = 0
    while i < n:
        if arr[i] == ' ':
            i += 1
            continue

        j = i
        while j < n and arr[j] != ' ':
            j += 1
        j -= 1

        sub_len = j - i + 1
        for k in range(sub_len // 2):
            arr[i + k], arr[i - k + sub_len - 1] = arr[i - k + sub_len - 1], arr[i + k]
        i = j + 1

    return arr


arr = ['a', 'b', 'c', ' ', ' ', 'd', 'e', ' ', 'f']
reverse(arr)
print(arr)
