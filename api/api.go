package api

const (
	RESPONSE_CODE_SUCCESS     string = "1"
	RESPONSE_CODE_FAIL        string = "0"
	RESPONSE_STATUS_CANCEL    string = "W"
	RESPONSE_STATUS_CANCEL_OK string = "R"
	RESPONSE_STATUS_CANCEL_N  string = "N"
)

type ResponseStatus string

const (
	PARAM_ILLEGAL          ResponseStatus = "参数非法"
	ILLEGAL_SIGN           ResponseStatus = "非法签名"
	NOTIFY_OD_FAIL         ResponseStatus = "回调处理失败"
	PARAM_NOT_NULL         ResponseStatus = "参数不能为空"
	REQUEST_ILLEGAL        ResponseStatus = "非法请求"
	TEL_IS_NULL            ResponseStatus = "手机号为空"
	TEL_EXIST              ResponseStatus = "手机号已存在"
	TEL_NOT_EXIST          ResponseStatus = "手机号不存在"
	TEL_NOT_BIND_OLD       ResponseStatus = "手机号不是目前绑定的"
	SMS_AUTHKEY_NOT_FOUND  ResponseStatus = "短信安全KEY不存在"
	SEND_SMS_FREQUENTLY    ResponseStatus = "短信发送频繁"
	USER_EXIST             ResponseStatus = "用户已经存在"
	USER_NOT_EXIST         ResponseStatus = "用户不存在"
	NICK_NAME_EXIST        ResponseStatus = "昵称已存在"
	NICK_NAME_NOT_EXIST    ResponseStatus = "昵称不存在"
	EMAIL_NOT_EXIST        ResponseStatus = "邮箱不存在"
	EMAIL_EXIST            ResponseStatus = "邮箱已存在"
	USER_NAME_EXIST        ResponseStatus = "用户名已存在"
	USER_NOT_INFO          ResponseStatus = "没有找到用户信息"
	CHECK_CODE_ERROR       ResponseStatus = "验证码错误"
	CHECK_CODE_INVALID     ResponseStatus = "验证码无效"
	MATCH_NOT_FOUND        ResponseStatus = "无法找到比赛信息"
	MATCH_NOT_CAN_BUY      ResponseStatus = "比赛未开售"
	MATCH_PURCHASED        ResponseStatus = "比赛已购买"
	PSSSWORD_ERROR         ResponseStatus = "密码错误"
	SEND_SMS_FAIL          ResponseStatus = "短信发送失败"
	WEIXIN_ERROR           ResponseStatus = "微信支付错误"
	WEIXIN_PARAMS_ALL_NULL ResponseStatus = "微信订单号与商户订单号其中一个必须有值"
	ACTIVITY_JOIN          ResponseStatus = "活动已经参加"
	ACTIVITY_FINISH        ResponseStatus = "活动已结束"
	ACCOUNT_NORMAL         ResponseStatus = "账号正常"
	ACCOUNT_STEAL          ResponseStatus = "账号异常"
	DIRECT_SUCCESS         ResponseStatus = "支付成功"
	SYS_EXCEPTION          ResponseStatus = "系统异常"
	PRICE_EXCEPTION        ResponseStatus = "金额为零"
	ORDER_NOT_EXIST        ResponseStatus = "订单不存在"
	ORDER_CANCELED         ResponseStatus = "订单已取消"
	PRODUCT_NOT_EXIST      ResponseStatus = "产品不存在"
	NOT_A_MEMBER           ResponseStatus = "未加入新英特权会员"
	NOT_VIP                ResponseStatus = "非会员"
	REQUEST_ERR            ResponseStatus = "请求异常"
)

type LocalResponse struct {
	ResCode   string         `json:"resCode"`
	ResStatus ResponseStatus `json:"resStatus"`
	Data      interface{}    `json:"data""`
	Msg       string         `json:"msg"`
	Extram    string         `json:"extram"`
}

const (
	APPRESPONSE_CODE_SUCCESS string = "200"
	APPRESPONSE_CODE_FAIL    string = "0"
)

type AppResponse struct {
	ResCode       string         `json:"resCode"`
	ResStatus     ResponseStatus `json:"resMessage"`
	RetData       interface{}    `json:"retData"`
	RequsetAction string         `json:"requsetAction"`
	UserToken     string         `json:"userToken"`
	ServerTime    string         `json:"serverTime"`
}
