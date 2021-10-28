# coding=utf-8

import sys
import math
import random
import pickle
import json
from collections import defaultdict
import redis

#r = redis.Redis(host='192.168.4.20',port=6379,password=888888, db=4)
usertagR = redis.Redis(host='coop.redis.ssports.com',port=6379,password='Redis42017sSports', db=4)
showliR = redis.Redis(host='activity.redis.ssports.com',port=6379,password='Redis42017sSports', db=4)

def push_user_tags(data):
    tagdict = defaultdict(dict)
    with open('user_tag_sample.txt') as fp:
        for line in fp:
            uid,tag,score = line.strip().split('\t')
            #if uid not in data:
                #continue
            tagdict[uid][tag] = score
    for u in tagdict:
        utagmap = tagdict[u]
        tagstr = json.dumps(utagmap)
        if len(utagmap) > 0:
            print(utagmap)
            print(u,tagstr)
            usertagR.set("usertags:"+u,tagstr)

def push_showclicklist(data):
    with open('uid_vid.txt') as fp:
        for line in fp:
            dt,uid,vid = line.strip().split('\t')
            #if uid not in data:
                #continue
            if dt < '2019-05-01':
                showliR.lpush('showli:'+uid, vid)
                showliR.lpush('quick_showli:'+uid, vid)
            else:
                showliR.lpush('clickli:'+uid, vid)
            data.add(uid)


def push_uuidmap():
    data = set()
    i = 0
    with open('uuid_map.txt') as fp:
        for line in fp:
            uuid,uid = line.strip().split('\t')
            data.add(uid)
            usertagR.set('uuidmap:'+uuid,uid)
            i += 1
            #if i > 1000:
                #break
    return data

if __name__ == '__main__':
    data = push_uuidmap()
    push_showclicklist(data)
    push_user_tags(data)
    print("get from redis")
    for uid in data:
        print(uid)
        #print(r.lrange('showli:'+uid,0,-1))
        #print(r.lrange('clickli:'+uid,0,-1))
        #print(r.lrange('quick_showli:'+uid,0,-1))
        jstr = usertagR.get('usertags:'+uid)
        if jstr is not None:
            print(json.loads(jstr))
