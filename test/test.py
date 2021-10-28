#!/usr/bin/python
# coding:utf-8

import sys
import random
from collections import defaultdict
import time

random.seed(time.time())

article_dict = defaultdict(dict)
idlist = []
with open("../data_pyscripts/vod_trans.txt") as fp:
    for line in fp:
        dt,uid,wtime,timelen,articleid,vc2title,tag_name,largeclass,mediumclass,smallclass = line.strip().split('\t')
        if articleid in article_dict:
            continue
        tag_name = tag_name.replace('[','').replace(']','').replace('"','').split(",")
        article_dict[articleid]["tags"] = tag_name
        article_dict[articleid]["large"] = largeclass
        article_dict[articleid]["medium"] = mediumclass
        article_dict[articleid]["small"] = smallclass
        idlist.append(articleid)

def prepare_list():
    def helper():
        reslist = []
        i = 0
        while i < 50:
            i+=1
            reslist.append(idlist[random.randrange(0,len(idlist))])
        return reslist
    newlist = helper()
    reclist = helper()
    cflist = helper()
    return (newlist,reclist,cflist)


def test_1(alllist):
    stgy = ("up","down","home")[random.randrange(0,3)]
    fixed = {0:idlist[random.randrange(0,len(idlist))]}
    fixed[3] = idlist[random.randrange(0,len(idlist))]
    fixed[7] = idlist[random.randrange(0,len(idlist))]

    newlist,reclist,cflist = alllist
    i = 0

    prob = (0.5,0.75,1)
    minidx = [0,0,0]
    nowidx = [0,0,0]
    reslist = []

    while 1:
        i = 0
        randf = random.random()
        for _ in alllist:
            if randf > prob[i]:
                i += 1
                continue
            else:
                break
        # first choose a list
        tmplist = alllist[i]
        nowi = nowidx[i]
        mini = minidx[i]
        if nowi == len(tmplist): # 本列表已全部选完
            # 此处应检查badcase,即所有列表中均无法选出符合条件的内容, 可以设置一个map, 删除列表索引, 一旦选择成功, 需要重置该map
            continue

        while 1:
            if tmplist[nowi] in reslist:
                nowi += 1
            else:
                break
        if nowi == len(tmplist): # 本列表已全部选完
            continue
        # secend choose a item
        item = tmplist[nowi]

        # 检查逻辑 暂时改为以0.2的概率选择是否接受
        if 0.1 < random.random():
            # 不符合条件, 不会被选
            nowidx[i] = nowi+1 # 下一次从下一个位置开始选
            continue
        else:
            # 关键点, 重启整个循环
            reslist.append(tmplist[nowi])
            nowidx[i] = minidx[i]  # 既然已经选择了该列表中最靠前的符合条件的内容,  下一次选择条件会有所不同, 所以又从最小位置开始选

        if len(reslist) == 20:
            return reslist  # over

if __name__ == '__main__':
    time1 = time.time()
    loops = 1000000
    for i in range(loops):
        test_1(prepare_list())
    time2 = time.time()
    print((time2-time1)/loops)
