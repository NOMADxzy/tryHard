
# 4 8
# cut 15
# cute 10
# but 6
# execute 3

n, m = map(int, input().split())
words = []

dif_set = set()
ans = []
for _ in range(n):
    v1, v2 = input().split()
    words.append([list(v1), int(v2)])
    for s in list(v1):
        if s not in dif_set:
            dif_set.add(s)
            ans.append(s)

ans.sort()

while len(dif_set)<m:
    cnts = {}
    for word_list, k in words:
        for i in range(len(word_list)-1):
            key_str = str(word_list[i]) + "_" + str(word_list[i+1])
            if key_str not in cnts:
                cnts[key_str] = [word_list[i], 0]
            cnts[key_str][1] += k

    best_str, best_first_str, max_cnt = "", "", 0
    for key_str, v in cnts.items():
        first_str, cnt = v
        if cnt > max_cnt or cnt == max_cnt and (len(key_str)<len(best_str) or len(key_str)==len(best_str) and (len(first_str)<len(best_first_str) or len(first_str)==len(best_first_str) and first_str<best_first_str)):
            best_str, best_first_str, max_cnt = key_str, first_str, cnt
    if best_str == "":
        break

    else:
        true_best_str = best_str.replace("_", "")
        for j in range(len(words)):
            word_list, _ = words[j]
            new_word_list = []
            i = 0
            while i+1<len(word_list):
                key_str = str(word_list[i]) + "_" + str(word_list[i + 1])
                if key_str == best_str:
                    words[j][0] = word_list[:i] + [true_best_str] + word_list[i+2:]
                    new_word_list.append(true_best_str)
                    i += 2
                else:
                    new_word_list.append(word_list[i])
                    i += 1
            if i==len(word_list)-1:
                new_word_list.append(word_list[-1])
            words[j][0] = new_word_list
        if best_str not in dif_set:
            dif_set.add(true_best_str)
            ans.append(true_best_str)

print("\n".join(ans))

