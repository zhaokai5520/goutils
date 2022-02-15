package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhaokai5520/goutils/errorx"
)

const defaultCode = 1001

// const okCode = 10000

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//正确返回的json
func Success(w http.ResponseWriter, v interface{}) {
	httpx.OkJson(w, Response{
		Code: http.StatusOK,
		Msg:  "sucess",
		Data: v,
	})
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

// defaultCode返回码，自定义返回信息
func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

// 根据状态码返回定义错误
func DefCodeError(code int) error {
	return &CodeError{Code: code, Msg: errorx.GetMsg(code)}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
