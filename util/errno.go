package util

type Errno struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e Errno) New(code int, msg string, data interface{}) *Errno {
	return &Errno{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

var (
	SUCCESS = Errno{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}

	FAILURE = Errno{
		Code: 500,
		Msg:  "fail",
		Data: nil,
	}

)
