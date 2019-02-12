package captcha

import (
	"testing"
	"fmt"
)

// Get phone captcha test.
func TestGetPhoneCaptcha(t *testing.T) {
	captcha, err := GetPhoneCaptcha()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("phone captcha is: ", captcha)
}

// Get email captcha test.
func TestGetEmailCaptcha(t *testing.T) {
	captcha, err := GetEmailCaptcha()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("email captcha is: ", captcha)
}