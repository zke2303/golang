package errcode

type ErrCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrCode) Error() string {
	return e.Message
}

// 通用错误码
var (
	OK                  = ErrCode{0, "success"}
	ParamsErr           = ErrCode{10001, "parameter error"}
	ParamsMissing       = ErrCode{10002, "parameter missing"}
	IllegalParams       = ErrCode{10003, "Illegal Parameter Format"}
	NotFound            = ErrCode{10004, "Data not fount"}
	Exists              = ErrCode{10005, "The data already exists."}
	Forbidden           = ErrCode{10006, "Forbidden"}
	Unauthorized        = ErrCode{10007, "Not Logged in"}
	InternalServerError = ErrCode{10008, "InternalServerError"}
)
