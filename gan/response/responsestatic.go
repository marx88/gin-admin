package response

import "github.com/gin-gonic/gin"

// New 实例化
func New() *Response {
	return &Response{gin.H{}}
}

// SuccessMsgData 成功
func SuccessMsgData(c *gin.Context, msg string, data interface{}) {
	New().SuccessMsgData(c, msg, data)
}

// SuccessMsg 成功
func SuccessMsg(c *gin.Context, msg string) {
	New().SuccessMsg(c, msg)
}

// SuccessData 成功
func SuccessData(c *gin.Context, data interface{}) {
	New().SuccessData(c, data)
}

// Success 成功
func Success(c *gin.Context) {
	New().Success(c)
}

// SuccessPage 成功分页
func SuccessPage(c *gin.Context, total int64, rows interface{}) {
	New().SuccessPage(c, total, rows)
}

// ErrorMsgData 失败
func ErrorMsgData(c *gin.Context, msg string, data interface{}) {
	New().ErrorMsgData(c, msg, data)
}

// ErrorMsg 失败
func ErrorMsg(c *gin.Context, msg string) {
	New().ErrorMsg(c, msg)
}

// ErrorData 失败
func ErrorData(c *gin.Context, data interface{}) {
	New().ErrorData(c, data)
}

// Error 失败
func Error(c *gin.Context) {
	New().Error(c)
}

// ErrorPage 分页失败
func ErrorPage(c *gin.Context, msg string) {
	New().ErrorPage(c, msg)
}
