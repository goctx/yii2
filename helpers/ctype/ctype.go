package ctype

import "regexp"

// 检测是否纯字母或数字
func AlNum(text string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", text)
	return match
}
