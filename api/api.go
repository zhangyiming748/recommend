package api

const (
	RESPONSE_CODE_SUCCESS     string = "1"
	RESPONSE_CODE_FAIL        string = "0"
	RESPONSE_STATUS_CANCEL    string = "W"
	RESPONSE_STATUS_CANCEL_OK string = "R"
	RESPONSE_STATUS_CANCEL_N  string = "N"
)

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
	APPRESPONSE_CODE_SUCCESS string = "200"
	APPRESPONSE_CODE_FAIL    string = "0"
)