package network

type Code int

// 定义返回Code.
const (
	FAILED  Code = -1   // 成功
	SUCCESS Code = iota // 失败
)
