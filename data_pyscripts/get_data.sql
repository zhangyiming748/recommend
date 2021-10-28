select max(dt) as dt,bm.uid as uid,vid from
    (
    select dt,user_id as uid, cast(article_id as bigint) as vid from data_model.article_daily_collect
    where dt >= date_sub('2019-08-21',90)
    union all
    select dt,user_id as uid, cast(vid as bigint) from data_model.vod_daily_collect
    where dt >= date_sub('2019-08-21',90)
    union all
    select dt,uid, cast(mid as bigint) from data_model.live_daily
    where dt >= date_sub('2019-08-21',90)
        ) tmp
    join ssports_shop.bb_member bm
    on tmp.uid = bm.uid
where bm.uid != ''
group by bm.uid,vid
    ;
