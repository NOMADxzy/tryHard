# 模拟文字编辑器
opts = input().split('|')
cur_lines = []
error = False
for cmd in opts:
    splits = cmd.split(maxsplit=2)
    idx, opt = int(splits[0]), splits[1]
    if opt == 'd':
        if idx - 1 < len(cur_lines):
            error = True
            break
        cur_lines = cur_lines[:idx - 1] + cur_lines[idx:]
    else:
        content = splits[2]
        if opt == 'a':
            if idx-1 < len(cur_lines):
                error = True
                break
            cur_lines = cur_lines[:idx] + [content] + cur_lines[idx:]
        elif opt == 'i':
            if idx < len(cur_lines):
                error = True
                break
            cur_lines = cur_lines[:idx-1] + [content] + cur_lines[idx-1:]
        else:
            if idx-1 < len(cur_lines):
                error = True
                break
            cur_lines[idx-1] = content

if error:
    print("error")
else:
    for line in cur_lines:
        print(line)