package captcha

import "errors"

// Get phone captcha.
func GetPhoneCaptcha() (string, error) {
	return "654321", nil
}

// Get email captcha.
func GetEmailCaptcha() (string, error) {
	return "123456", nil
}


// Get captcha.
func GetCaptcha(captchaStyle, featureStyle int) (string, error) {
	var captcha string
	switch captchaStyle {
	case 0:

	case 1:

	default:
		return captcha, errors.New("undefined captcha style")
	}

	return captcha, nil
}

