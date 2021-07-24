package middle

import (
	"time"
)

// R -
type R struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

// Status -
type Status struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Unix int64  `json:"unix"`
}

// Response -
func response(data interface{}, err error) R {
	t := time.Now()
	return R{
		Data: data,
		Status: Status{
			Msg:  err.Error(),
			Unix: t.Unix(),
		},
	}
}

// Error -
func Error(err error) R {
	return response(nil, err)
}

// Resp -
func Resp(data interface{}) R {
	return response(data, nil)
}
