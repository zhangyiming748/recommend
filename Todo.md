# 还有个内容, 顶部模板也需要解析
2019/06/25 20:47:46 template.go:112 ERROR	 cannot get articleDetail from redis: 2021488 err: <nil>
2019/06/25 20:47:46 template.go:112 ERROR	 cannot get articleDetail from redis: 2021159 err: <nil>

/*
公共参数 uid也是公共参数, abteststgy, algostgy 公共参数
传入的是过滤的filterMap[articleid] 绝对不出的, 这可能为空
1. new 发布时间两天内, 按发布时间倒序排列, 取50条, 如果不够, 有多少给多少
2. rec 这块需要和用户画像结合,rankserver 发布时间大于2天, 2天前的, 按时间倒序排序, top50 后续这块就top500
3. usercf 取50个相近用户, 取50条, 如果不够, 有多少给多少

返回 RecommendList
3.
*/


# 有一个很矛盾的点 首先, 我必须从redis中获取articletags, 实时画像词需要
# 其次, 我又不能依赖redis, 因为内容下线的话, 我不知道, 我必须从ES中获取内容, 本地不敢缓存内容,这样才能确保出的内容, 是确定可出的内容.

# 数据结构
## 曝光列表 showli_uid lpush入队,最新的在最前
## 点击列表 clickli_uid lpush入队,最新的在最前
## 相似用户列表 userSimilar_uid json字符串,需要首先序列化 结构为 [(uid,score),{uid,score}...], 按照分数排序,高分在前

# recommend-url
## 1.配置域名和端口, 先直接返回静态文件
/.vid.txt  
random articleid
get from redis
替换模板
## 2.解析请求, 是上滑还是下拉刷新, 获取userid

## 3.使用userid, 读取redis, 拿到用户的已曝光过滤列表, 以及用户的最近几个点击, 形式为 两个articleid的队列
uid_showli uid_clickli
redis list结构
llen 
需要考虑左端进还是右端进的性能 lpush rpush, 允许重复
读取使用lrange[0:500] lrange[0:5]
xzhao
## 4.两个article队列合并, 点击的article作为已曝光的内容的补充, 同样过滤. 模板中内容同样作为已曝光过滤内容.
xzhao
## 5.获取用户最近(5个)的点击队列的大中小项词和tag词, 计算实时画像词分数, 并同时从redis中读取历史画像词
article_taginfo
{
tags:
largeclass
mediumclass
smallclass
datapublish
}
xzhao
## 6.请求下游服务, 请求中带上已曝光队列和画像词分数(无论请求usercf, 新内容, 或是推荐排序内容)
pb zhao 学习一下
## 7.按照画像词, 按照固定比率打散, 得到最终articleid的列表


## 8.读取redis, 获得每个articleid的模板,填充接口模板.
rf已经完成80%

## 9.打印日志
返回前需要需要将下发列表存入redis, 以作为快速过滤列表 lpush quick_showli_uid [items..]
快速过滤列表仅仅存储40条, 主要是防止以下可能:
用户请求了20条内容, 仅观看当前屏幕的5条, 然后立刻下拉刷新, 得到新的十条内容. 此时若用户连续翻页, 则可能会发现头十条和未曝光的十条中, 存在重复内容.

# rank-server
## 1.获得recommend-url的pb请求
请求中包含
uid,
设备id,
用户画像分(暂时不用),
需过滤的article列表(不出的article)
以及召回条件(暂时按照时间倒序排序, 出500条即可)
## 2.按照要求, 请求ES

## 3.返回结果给recommend-url
返回列表为
[
{
articleid:
article tags:
article publish time
article 大中小项
},
...
]
