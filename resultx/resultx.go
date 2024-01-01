package resultx

import (
	"context"

	"github.com/amorist/common/errcode"
)

const (
	defaultCode    = 0
	defaultErrCode = 1
	defaultMsg     = "请求成功"
)

// ResultResponse .
type ResultResponse struct {
	ID   string      `json:"id,omitempty"` // TraceId
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewCodeResponse .
func NewCodeResponse(code int, msg string, data interface{}) *ResultResponse {
	return &ResultResponse{Code: code, Msg: msg, Data: data}
}

// NewResponse .
func NewResponse() *ResultResponse {
	return &ResultResponse{Code: defaultCode, Msg: defaultMsg, Data: nil}
}

// Error .
func (r *ResultResponse) Error(err error) *ResultResponse {
	switch err.(type) {
	case *errcode.Error:
		e := err.(*errcode.Error)
		r.Code = e.Code
		r.Msg = e.Msg
	default:
		r.DefaultError(err.Error())
	}
	return r
}

// CodeError .
func (r *ResultResponse) CodeError(code int, msg string) *ResultResponse {
	r.Code = code
	r.Msg = msg
	return r
}

// DefaultError .
func (r *ResultResponse) DefaultError(msg string) *ResultResponse {
	r.CodeError(defaultErrCode, msg)
	return r
}

// WithTraceID .
func (r *ResultResponse) WithTraceID(ctx context.Context) *ResultResponse {
	// t, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	// if ok {
	// 	r.ID = t.TraceId()
	// }
	return r
}
