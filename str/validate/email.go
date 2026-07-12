package validate

import "regexp"

// IsEmail 简单验证邮箱格式
func IsEmail(email string) bool {
	const pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// IsPhone 验证中国大陆手机号 (11位，1开头)
func IsPhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
