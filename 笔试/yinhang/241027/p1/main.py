# 给一个num，找出三个素数之和等于num的所有不同组合数量

# 100
class Solution:
    """
    【编程】含通配符字符串查找
    :param mainstr: 主字符串
    :param substr: 子字符串
    :return: 次数
    """

    def getTimes(self, mainstr, substr):
        ans = 0
        tmp = []
        for i in range(len(substr)):
            if substr[i] == '?':
                tmp.append(i)
        i1, i2 = tmp
        l1, l2, l3 = i1, i2 - i1 - 1, len(substr) - 1 - i2
        list1, list2, list3 = [], [], []

        for i in range(len(mainstr)):
            if i + l1 <= len(mainstr) and mainstr[i:i + l1] == substr[0:i1]:
                list1.append(i)
            if i + l2 <= len(mainstr) and mainstr[i:i + l2] == substr[i1 + 1:i2]:
                list2.append(i)
            if i + l3 <= len(mainstr) and mainstr[i:i + l3] == substr[i2 + 1:]:
                list3.append(i)
        # return list2
        for i in list1:
            for j in list2:
                if i + l1 > j:
                    continue
                for idx in range(len(list3)):
                    k = list3[idx]
                    if j + l2 > k:
                        continue
                    else:
                        # ans.append((i,j,len(list3) - idx-idx))

                        ans += len(list3) - idx
                        break

        return ans