package handlers

import "bilibili-rear-end/network"

type Resc int

const (
	parameterFailure Resc = iota		// 参数错误

	nickNameWrongFormate	// 昵称格式错误

	userNameWrongFormate	// 用户名格式错误

	phoneWrongFormate	// 手机号码格式错误
	emailWrongFormate	// 邮箱地址格式错误
	userOrPasswordWrong	// 用户名/密码错误

	registerFailure	// 注册失败
	signInFailure	// 登录失败
	signOutFailure // 登出失败
)

func (this Resc) Message() string {
	switch this {
	case parameterFailure:
		return "参数异常"
	case nickNameWrongFormate:
		return "用户名格式错误"
	case phoneWrongFormate:
		return "手机号码格式错误"
	case emailWrongFormate:
		return "邮箱格式错误"
	case userOrPasswordWrong:
		return "用户名/密码错误"
	case registerFailure:
		return "注册失败"
	case signInFailure:
		return "登录失败"
	case signOutFailure:
		return "登出失败"
	}
	return ""
}

func (this Resc) DisplayMsg() string {
	switch this {
	case parameterFailure:
		return "参数异常"
	case nickNameWrongFormate:
		return "用户名格式错误"
	case phoneWrongFormate:
		return "手机号码格式错误"
	case emailWrongFormate:
		return "邮箱格式错误"
	case userOrPasswordWrong:
		return "用户名/密码错误"
	case registerFailure:
		return "注册失败"
	case signInFailure:
		return "登录失败"
	case signOutFailure:
		return "登出失败"
	}

	return ""
}

func (this Resc) Code() network.Code {
	switch this {

	}

	return network.FAILURE
}