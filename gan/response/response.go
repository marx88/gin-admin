package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// NameCode code名称
	NameCode = "code"

	// NameMsg msg名称
	NameMsg = "msg"

	// NameData data名称
	NameData = "data"

	// NameTotal total名称
	NameTotal = "total"

	// NameRows rows名称
	NameRows = "rows"

	// DefaultMsgSuccess 默认成功消息
	DefaultMsgSuccess = "操作成功"

	// DefaultMsgError 默认失败消息
	DefaultMsgError = "操作失败"

	// CodeSuccess 成功状态码
	CodeSuccess = 200

	// CodeError 失败状态码
	CodeError = 500
)

// Response 响应
type Response struct {
	h gin.H
}

func (r *Response) code(code int) {
	r.Put(NameCode, code)
}

func (r *Response) msg(msg string) {
	r.Put(NameMsg, msg)
}

func (r *Response) data(data interface{}) {
	if data != nil {
		r.Put(NameData, data)
	}
}

// Put 把键值对存入h
func (r *Response) Put(key string, value interface{}) {
	r.h[key] = value
}

// Send 发送response
func (r *Response) Send(c *gin.Context) {
	c.JSON(http.StatusOK, r.h)
}

// SuccessMsgData 成功
func (r *Response) SuccessMsgData(c *gin.Context, msg string, data interface{}) {
	r.code(CodeSuccess)
	r.msg(msg)
	r.data(data)
	r.Send(c)
}

// SuccessMsg 成功
func (r *Response) SuccessMsg(c *gin.Context, msg string) {
	r.SuccessMsgData(c, msg, nil)
}

// SuccessData 成功
func (r *Response) SuccessData(c *gin.Context, data interface{}) {
	r.SuccessMsgData(c, DefaultMsgSuccess, data)
}

// Success 成功
func (r *Response) Success(c *gin.Context) {
	r.SuccessData(c, nil)
}

// SuccessPage 成功分页
func (r *Response) SuccessPage(c *gin.Context, total int, rows interface{}) {
	r.Put(NameTotal, total)
	r.Put(NameRows, rows)
	r.Success(c)
}

// ErrorMsgData 失败
func (r *Response) ErrorMsgData(c *gin.Context, msg string, data interface{}) {
	r.code(CodeError)
	r.msg(msg)
	r.data(data)
	r.Send(c)
}

// ErrorMsg 失败
func (r *Response) ErrorMsg(c *gin.Context, msg string) {
	r.ErrorMsgData(c, msg, nil)
}

// ErrorData 失败
func (r *Response) ErrorData(c *gin.Context, data interface{}) {
	r.ErrorMsgData(c, DefaultMsgError, data)
}

// Error 失败
func (r *Response) Error(c *gin.Context) {
	r.ErrorData(c, nil)
}

// ErrorPage 分页失败
func (r *Response) ErrorPage(c *gin.Context, total int, rows interface{}) {
	r.Put(NameTotal, total)
	r.Put(NameRows, rows)
	r.Error(c)
}
