package helper

import "regexp"

// IsChinese 是否是中文
// 包含简/繁体 扩大生僻字范围
func IsChinese(s string) bool {
	// 查找连续的汉字
	reg := regexp.MustCompile(`^[\p{Han}]+$`)
	matched := reg.Match([]byte(s))
	return matched
}
