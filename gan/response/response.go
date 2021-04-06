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
	r.msg(msg)
	r.data(data)
	r.Success(c)
}

// SuccessMsg 成功
func (r *Response) SuccessMsg(c *gin.Context, msg string) {
	r.msg(msg)
	r.Success(c)
}

// SuccessData 成功
func (r *Response) SuccessData(c *gin.Context, data interface{}) {
	r.msg(DefaultMsgSuccess)
	r.data(data)
	r.Success(c)
}

// Success 成功
func (r *Response) Success(c *gin.Context) {
	r.code(CodeSuccess)
	r.Send(c)
}

// SuccessPage 成功分页
func (r *Response) SuccessPage(c *gin.Context, total int64, rows interface{}) {
	r.Put(NameTotal, total)
	r.Put(NameRows, rows)
	r.msg(DefaultMsgSuccess)
	r.Success(c)
}

// ErrorMsgData 失败
func (r *Response) ErrorMsgData(c *gin.Context, msg string, data interface{}) {
	r.msg(msg)
	r.data(data)
	r.Error(c)
}

// ErrorMsg 失败
func (r *Response) ErrorMsg(c *gin.Context, msg string) {
	r.msg(msg)
	r.Error(c)
}

// ErrorData 失败
func (r *Response) ErrorData(c *gin.Context, data interface{}) {
	r.msg(DefaultMsgError)
	r.data(data)
	r.Error(c)
}

// Error 失败
func (r *Response) Error(c *gin.Context) {
	r.code(CodeError)
	r.Send(c)
}

// ErrorPage 分页失败
func (r *Response) ErrorPage(c *gin.Context, msg string) {
	r.Put(NameTotal, 0)
	r.Put(NameRows, nil)
	r.msg(msg)
	r.Error(c)
}
