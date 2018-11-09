package defs

//直接提取url中json:<变量名>的value
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

//嵌套上边的结构体并添加HttpSC值代表httpCode值，作为错误处理的结构体
type ErrorReponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorReponse{HttpSC: 400, Error: Err{Error: "Request body is not correct.", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrorReponse{HttpSC: 401, Error: Err{Error: "User authentication failed. ", ErrorCode: "002"}}
	ErrorDBError                = ErrorReponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorReponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
