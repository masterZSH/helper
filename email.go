package helper

import "regexp"

// IsEmail 是否是邮箱
func IsEmail(email string) bool {
	matched, err := regexp.Match(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, []byte(email))
	if err != nil {
		return false
	}
	return matched
}
