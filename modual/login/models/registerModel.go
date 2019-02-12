package models

import (
	database "bilibili-rear-end/database/mysql"
	"errors"
	"time"
	"fmt"
)

const defaultTimesFormate = "2006-01-02 15:04:05"
// User signIn info.
type User struct {
	account string // login account
	password string // 密码(暂时明文)
	createdTime string // register time
	lastSignInTime string // last signIn time
	lastSignOutTime string // last sign out time
}

// Create a new user.
func NewUser(nickname, password, account string, registerStyle int) error {

	var commonContact string
	switch registerStyle {
	case 0:	// mobile
		commonContact = "mobile"
	case 1:	// email
		commonContact = "email_address"
	}

	const sql = "INSERT INTO `users` (`account`, `password`, `created_time`) VALUES (?, ?, ?) SELECT `%s`, `nick_name` FROM `users_info` "

	result, err := database.MemberDB().Insert(fmt.Sprintf(sql, commonContact), account, password, time.Now().Format(defaultTimesFormate), account, nickname)
	if err != nil {
		return err
	}

	if result < 0 {
		return errors.New("Register failure.")
	}

	return nil
}

// Query user with account.
func QueryUser(account string) (*User, error) {
	const sql = "SELECT `account` FROM `users` WHERE `account` = ?"

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