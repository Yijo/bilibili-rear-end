package network

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	defaultTimesFormate = "2006-01-02 15:04:05"

	serverException = "服务器异常"
)

// Unified interface response.
type ResponseAdapter struct {
	status     int         // HTTP status code
	Code       Code        `json:"code"`                 // return code
	Data       interface{} `json:"data,omitempty"`       // return data
	Error      string      `json:"error,omitempty"`      // return information when an error occurs
	Message    string      `json:"message,omitempty"`    // return description，used to display to developer
	DisplayMsg string      `json:"displayMsg,omitempty"` // return description，used to display to the user
	Timestamp  string      `json:"timestamp"`            // return timestamp
}

// Response description.
type ResponseDesc interface {
	Message() string
	DisplayMsg() string
	Code() Code
}

// Request success.
func Success(httpStatus int, desc ResponseDesc, data interface{}) *ResponseAdapter {
	adapter := &ResponseAdapter{
		status:     httpStatus,
		Data:       data,
		Code:       desc.Code(),
		Message:    desc.Message(),
		DisplayMsg: desc.DisplayMsg(),
		Timestamp:  time.Now().Format(defaultTimesFormate),
	}

	if desc.DisplayMsg() == "" && desc.Message() == "" {
		adapter.Message = serverException
		adapter.DisplayMsg = serverException
	}

	return adapter
}

// Request failure.
func Failure(httpStatus int, desc ResponseDesc) *ResponseAdapter {
	adapter := &ResponseAdapter{
		status:     httpStatus,
		Code:       desc.Code(),
		Message:    desc.Message(),
		DisplayMsg: desc.DisplayMsg(),
		Timestamp:  time.Now().Format(defaultTimesFormate),
	}

	if desc.DisplayMsg() == "" && desc.Message() == "" {
		adapter.Message = serverException
		adapter.DisplayMsg = serverException
	}

	return adapter
}

// Write data.
func (adapter *ResponseAdapter) Response(c *gin.Context) {
	// for the time being, only return JSON first, do not judge Accept and ContentType
	c.JSON(adapter.status, adapter)
}


// Add message description.
func (adapter *ResponseAdapter)AppendMessage(str string) *ResponseAdapter {
	if adapter.Message == "" {
		return adapter.SetMessage(str)
	} else {
		adapter.Message = adapter.Message + ", " + str
		return adapter
	}
}

// Set message description.
func (adapter *ResponseAdapter)SetMessage(str string) *ResponseAdapter {
	adapter.Message = str
	return adapter
}