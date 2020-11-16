package helper

import (
	"bytes"
	"regexp"
)

// IsPhoneNumber  是否是手机号
func IsPhoneNumber(phoneNumber string) bool {
	matched, err := regexp.Match(`1[3-9][0-9]{9}`, []byte(phoneNumber))
	if err != nil {
		return false
	}
	return matched
}

// GetPhoneCipherText 加密手机号
// "13112312312" => "131****1234"
func GetPhoneCipherText(phoneNumber string) string {
	if !IsPhoneNumber(phoneNumber) {
		return ""
	}
	var res bytes.Buffer
	res.WriteString(phoneNumber[:3])
	res.WriteString("****")
	res.WriteString(phoneNumber[7:])
	return res.String()
}


