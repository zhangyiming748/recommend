# coding=utf-8

import sys
import math
import random
import pickle
import json
from collections import defaultdict
import redis

r = redis.Redis(host='192.168.4.20',port=6379,password=888888, db=1)
similarR = redis.Redis(host='192.168.4.20',port=6379,password=888888, db=4)

test_usr_items = defaultdict(set)

with open('uid_vid.txt') as fp:
    for line in fp:
        dt,uid,vid = line.strip().split('\t')
        if uid in ('NULL',''):
            continue
        if dt < '2019-05-01':
            continue
        else:
            test_usr_items[uid].add(vid)
#print(similarR.get("userSimilar_a1580267722c4263998458e78865f558"))
#print(similarR.get("userSimilar_"+"12050478")) 找不到这个key 这是因为 OOM command not allowed when used memory > 'maxmemory'.
#print(json.loads(similarR.get("userSimilar_"+"a1580267722c4263998458e78865f558")))
#sys.exit(0)

hasnorecommend=0

def GetRecomendation(user,K):
    # 推荐测试
    rank = defaultdict(float)
    jstr = similarR.get("userSimilar_"+user)
    #print(jstr)
    if jstr is None:
        return None
    wu = json.loads(jstr)
    #print(len(wu))
    for (v, wuv) in wu[0:K]:
        if v == user:
            continue
        similar_interacted_items = r.lrange('clickli_'+v,0,-1)
        for i in similar_interacted_items:
            rank[i] += float(wuv) * 1.0
    return rank

def Recall(K):
    hit = 0
    allx = 0
    for user in test_usr_items.keys():
        tu = r.lrange('clickli_'+user,0,-1)
        tu = set(tu)
        #print(user,tu)
        rank = GetRecomendation(user,K)
        if rank is None:
            global hasnorecommend
            hasnorecommend+=1
            continue
        for item, pui in rank.items():
            if item in tu:
                hit += 1
        allx += len(tu)
    if allx == 0:
        return 0
    return hit / (allx * 1.0)

def Precision(K):
    hit = 0
    allx = 0
    for user in test_usr_items.keys():
        tu = r.lrange('clickli_'+user,0,-1)
        tu = set(tu)
        rank = GetRecomendation(user,K)
        if rank is None:
            continue
        for item, pui in rank.items():
            if item in tu:
                hit += 1
        allx += len(rank)
    if allx == 0:
        return 0
    return hit / (allx * 1.0)

def Coverage(K):
    recommend_items = set()
    all_items = set()
    for user in test_usr_items.keys():
        for item in test_usr_items[user]:  #仅为测试集合中流行度
            all_items.add(item)
        rank = GetRecomendation(user,K)
        for item, pui in rank.items():
            recommend_items.add(item)
    return len(recommend_items) / (len(all_items)*1.0)

def Popularity(K):
    item_popularity = defaultdict(int)
    for user, items in test_usr_items.items():
        for item in items:
            item_popularity[item] += 1
    ret = 0
    n = 0
    for user in test_usr_items.keys():
        rank = GetRecomendation(user,K)
        for item, pui in rank.items():
            ret += math.log(1 + item_popularity[item])
            n += 1
    if n == 0:
        return 0
    ret /= n*1.0
    return ret

if __name__ == '__main__':
    import numpy as np
    print(Recall(20))
    print(len(test_usr_items), hasnorecommend)
    #print(GetRecomendation("a1580267722c4263998458e78865f558",20))
    #for K in np.arange(1,300):
    for K in (1,5,10,20,30,40,50,80,100,200):
        recall = Recall(K)
        precision = Precision(K)
        coverage = Coverage(K)
        popularity = Popularity(K)
        ss = '\t'.join(['neighbours:',str(K),'recall:',str(recall), 'precision:',str(precision), 'coverage:',str(coverage), 'popularity:',str(popularity)])
        print(ss)
