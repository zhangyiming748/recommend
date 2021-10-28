select a.deviceid, uid from (
    select deviceid,count(distinct(uid)) as uidnum from (
        select
            deviceid,
            uid
        from data_model.user_device_bind_day
        where dt >= '2019-06-05'
            and uid not in ('','NULL','null','0')
            and deviceid not in ('000000000000000')
            and plat = 'app'
            ) tmp
    group by deviceid
        ) a
    join (
    select
        deviceid,
        uid
    from data_model.user_device_bind_day
    where dt >= '2019-06-05'
        and uid not in ('','NULL','null','0')
        and deviceid not in ('000000000000000')
        ) b
    on a.deviceid = b.deviceid
where uidnum = 1 --对应多个uid的设备id不使用
group by a.deviceid,uid
    ;
