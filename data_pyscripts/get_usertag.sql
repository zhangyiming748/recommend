select
    uid,tag,score
from data_mining.vod_usr_positive_tags
where dt = '2019-05-30' and uid != ''
;
