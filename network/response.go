package network

import (
	"github.com/gin-gonic/gin"
	"time"
)

// 用于统一接口响应.
type ResponseAdapter struct {
	status     int         // HTTP状态码
	Code       Code        `json:"code"`                 // 返回码
	Data       interface{} `json:"data,omitempty"`       // 返回数据
	Error      string      `json:"error,omitempty"`      // 错误时返回信息
	Message    string      `json:"message,omitempty"`    // 返回描述，用于开发人员查看
	DisplayMsg string      `json:"displayMsg,omitempty"` // 返回描述，用于显示给用户查看
	Timestamp  string      `json:"timestamp"`            // 返回时间戳
}

// 响应描述.
type ResponseDesc interface {
	Message() string
	DisplayMsg() string
}

// 定义返回Code.
const (
	failed  int = -1   // 成功
	success int = iota // 失败
)

const defaultTimesFormate = "2018-11-02 15:04:05"

// 请求成功.
func Success(httpStatus int, code Code, desc ResponseDesc, data interface{}) *ResponseAdapter {
	return &ResponseAdapter{
		status:     httpStatus,
		Code:       code,
		Data:       data,
		Message:    desc.Message(),
		DisplayMsg: desc.DisplayMsg(),
		Timestamp:  time.Now().Format(defaultTimesFormate),
	}
}

// 请求失败.
func Failure(httpStatus int, desc ResponseDesc) *ResponseAdapter {
	return &ResponseAdapter{
		status:     httpStatus,
		Code:       FAILED,
		Message:    desc.Message(),
		DisplayMsg: desc.DisplayMsg(),
		Timestamp:  time.Now().Format(defaultTimesFormate),
	}
}

// 设置错误信息.
func (adapter *ResponseAdapter) AppendError(errMsg string) *ResponseAdapter {
	adapter.Error = errMsg
	return adapter
}

// 写入数据.
func (adapter *ResponseAdapter) Response(c *gin.Context) {
	// 暂时先只返回JSON, 不判断Accept和ContentType
	c.JSON(adapter.status, adapter)
}
