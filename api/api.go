package api
type LocalResponse struct {
	ResCode   string         `json:"resCode"`
	ResStatus ResponseStatus `json:"resStatus"`
	Data      interface{}    `json:"data"`
	Msg       string         `json:"msg"`
	Extram    string         `json:"extram"`
}

type ResponseStatus string
type AppResponse struct {
	ResCode       string         `json:"resCode"`
	ResStatus     ResponseStatus `json:"resMessage"`
	RetData       interface{}    `json:"retData"`
	RequsetAction string         `json:"requsetAction"`
	UserToken     string         `json:"userToken"`
	ServerTime    string         `json:"serverTime"`
}

const (
	REQUEST_ERR ResponseStatus = "请求异常"
	APPRESPONSE_CODE_SUCCESS string = "200"
	RESPONSE_CODE_FAIL        string = "0"
)
