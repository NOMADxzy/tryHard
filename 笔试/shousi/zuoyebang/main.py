# 只有长宽比当前书更小的书才能放在当前书上，求最大高度
book_list = [(5, 5), (4, 4), (4, 9), (3, 8), (2, 7), (1, 6)]

# (5,5),(4,9),(4,4),(3,8),(2,7),(1,6)

length_books = {}
length_list = []
for book in book_list:
    l, w = book
    if l not in length_books:
        length_books[l] = [w]
        length_list.append(l)
    else:
        length_books[l].append(w)

print(length_books)
print(length_list)

for l, books in length_books.items():
    books.sort()
    length_books[l] = books


def dfs(pos, limit):
    if pos == len(length_list):
        return 0
    l = length_list[pos]
    cur_books = length_books[l]
    i = len(cur_books) - 1
    while i >= 0 and cur_books[i] > limit:
        i -= 1
    if i == -1:
        return 0
    else:
        return 1 + dfs(pos + 1, cur_books[i])


ans = 0
for p in length_list:
    ans = max(ans, dfs(p, float('inf')))

print(ans)
