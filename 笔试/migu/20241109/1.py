songList = list(input().split(','))
word1, word2 = input(), input()
last_word = None
last_idx = -1
ans = len(songList)

for i in range(len(songList)):
    word = songList[i]
    if word != word1 and word != word2:
        continue
    if last_word is None:
        last_word = word
        last_idx = i
        continue
    else:
        if last_word != word:
            ans = min(ans, i - last_idx - 1)
        last_word = word
        last_idx = i

print(ans)
