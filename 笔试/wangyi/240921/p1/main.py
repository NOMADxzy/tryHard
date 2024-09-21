# 10个游戏，每行表示某人对10个游戏的评价，0表示没玩这个游戏，1表示玩了没投这个游戏，2表示玩了且投了这个游戏
# 每个人票数 = min(5, 玩过游戏数 - 2)，否则非法不计入
# 求最多票数的游戏

n = int(input())

votes = [0 for _ in range(10)]
for _ in range(n):
    detail = list(map(int, input().split()))
    vote_idxs = []
    total_play_cnt = 0
    for idx in range(10):
        if detail[idx] == 0:
            pass
        elif detail[idx] == 1:
            total_play_cnt += 1
        else:
            total_play_cnt += 1
            vote_idxs.append(idx)
    if total_play_cnt >= 3 and len(vote_idxs) <= min(5, total_play_cnt-2):
        for idx in vote_idxs:
            votes[idx] += 1

max_val = max(votes)
ans = []
# print(votes)
for i in range(10):
    if votes[i] == max_val:
        ans.append(i+1)

print(" ".join(map(str, ans)))