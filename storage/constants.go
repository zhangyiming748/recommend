package storage

const (
	// 已曝光列表
	REDIS_SHOWLIST_KEY      = "showli"
	REDIS_CLICKLIST_KEY     = "clickli"
	REDIS_QUICKSHOWLIST_KEY = "quick_showli"
	REDIS_USERSIMI_KEY      = "userSimilar"
	REDIS_UUIDMAP_KEY       = "uuidmap"
	REDIS_USERTAGS_KEY      = "usertags"
	REDIS_VERIFYCODE_KEY    = "VerificationCode"

	//hgetall articleDetail:1787866
	ARTICLEDETAIL_KEY = "articleDetail"

	//hgetall articleNfb:uid
	ARTICLENFB_KEY = "articleNfb"
	// get articleClick:aid
	ARTICLECLICK_KEY = "articleClick"
)
