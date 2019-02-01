package models

import "bilibili-rear-end/database"

// 用户登录信息
type User struct {
	account string // 用户名
	password string // 密码(暂时明文)
	createdTime string // 注册时间
	lastSignInTime string // 上次登录时间
	lastSignOutTime string // 上次退出时间
}

// 新增用户
func insertUser() {

}

// 查询用户.
func QueryUser(account string) (*User, error) {
	const sql = "SELECT `account` FROM `users_account` WHERE `account` = ?"

	result, err := database.MemberDB().FetchRow(sql, account)

	if err != nil {
		return nil, err
	}

	return newUser(result), nil
}


func newUser(user map[string]string) *User {
	return &User{
		account: user["account"],
	}
}