package storage

const (
	// 已曝光列表
	REDIS_SHOWLIST_KEY      = "showli"
	//点击列表
	REDIS_CLICKLIST_KEY     = "clickli"
	//快速过滤列表
	REDIS_QUICKSHOWLIST_KEY = "quick_showli"
	//相似用户
	REDIS_USERSIMI_KEY      = "userSimilar"
	REDIS_UUIDMAP_KEY       = "uuidmap"
	//用户标签
	REDIS_USERTAGS_KEY      = "usertags"
	REDIS_VERIFYCODE_KEY    = "VerificationCode"

	//hgetall articleDetail:1787866
	ARTICLEDETAIL_KEY = "articleDetail"

	//hgetall articleNfb:uid
	ARTICLENFB_KEY = "articleNfb"
	// get articleClick:aid
	ARTICLECLICK_KEY = "articleClick"
)
