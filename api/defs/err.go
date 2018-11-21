package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var(
	ErrorRequestBodyParse = ErrResponse{400, Err{Error:"Request body is not current.", ErrorCode:001}}
	ErrorAuthorNotUsed    = ErrResponse{403, Err{Error:"Author has not rules.", ErrorCode:002}}
)