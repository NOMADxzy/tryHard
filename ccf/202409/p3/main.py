import sys

n = int(input())
origin = []
for _ in range(n):
    origin.append(input())

rows = sys.stdin.readlines()
new_rows = []
for row in rows:
    if row[0] != "#":
        new_rows.append(row[:-1])
rows = new_rows
paddings_wrong = False

def get_num_detail(s):
    s = s[3:-3]
    l, r = s.split(" ")
    NN, MM = map(int, l[1:].split(","))
    nn, mm = map(int, r[1:].split(","))
    return NN, MM, nn, mm

def check_match(orign_start_idx, before_content):
    if orign_start_idx<0 or orign_start_idx+len(before_content)>len(origin):
        return False
    i = 0
    while i<len(before_content) and orign_start_idx+i<len(origin):
        if origin[orign_start_idx+i] != before_content[i]:
            break
        i += 1
    return i == len(before_content)

row_idx = 0
pre_nm = 0
acc_NN = 0
while row_idx<len(rows) and rows[row_idx][0] != "@":
    row_idx += 1

while row_idx<len(rows):
    if rows[row_idx][0] != "@":
        paddings_wrong = True
        break
    NN, MM, nn, mm = get_num_detail(rows[row_idx])
    NN += acc_NN-1
    nn -= 1
    if NN<pre_nm:
        paddings_wrong = True
        break
    row_idx += 1
    before_content, after_content = [], []
    while row_idx<len(rows):
        flag = rows[row_idx][0]
        if flag == "@":
            break
        elif flag != "+" and flag != "-" and flag != " ":
            paddings_wrong = True
            break
        else:
            content = rows[row_idx][1:]
            if flag == "-":
                before_content.append(content)
            elif flag == "+":
                after_content.append(content)
            else:
                before_content.append(content)
                after_content.append(content)
        row_idx += 1
    if paddings_wrong or len(before_content)!=MM or len(after_content)!=mm:
        break

    sig = 0
    for sigma in range(MM):
        if check_match(NN-sigma, before_content):
            origin = origin[:NN-sigma] + after_content + origin[NN-sigma+MM:]
            sig = sigma
            break
        elif check_match(NN+sigma, before_content):
            origin = origin[:NN + sigma] + after_content + origin[NN + sigma + MM:]
            sig = sigma
            break

    pre_nm = NN+MM+sig
    acc_NN += sig

if paddings_wrong:
    print("Patch is damaged")
else:
    print("\n".join(origin))