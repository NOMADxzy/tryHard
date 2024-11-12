s = "192168101"

ans = []


def dfs(i, pos, tmp_str):
    if pos == 4:
        if i == len(s):
            ans.append(tmp_str[1:])
            return
    if i >= len(s):
        return
    if s[i] == '0':
        dfs(i + 1, pos + 1, tmp_str + ".0")
    else:
        for j in range(i, min(i + 3, len(s))):
            if int(s[i:j+1]) <= 255:
                dfs(j + 1, pos + 1, tmp_str+"."+s[i:j + 1])
            else:
                break


dfs(0, 0, "")
print(ans)
