package utils

import (
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"regexp"
)

// DesensitizeEmail 用于脱敏邮箱地址
func DesensitizeEmail(email string) (string, error) {
	// 使用正则表达式匹配邮箱地址
	re := regexp.MustCompile(`(.)(.*)(@.*)`)
	matches := re.FindStringSubmatch(email)

	if len(matches) != 4 {
		return "", errors.New(consts.ErrEmailFormat)
	}

	// 只显示第一个字符和最后一个字符，中间用星号(*)代替
	desensitizedEmail := string(matches[1]) + "*****" + matches[3]

	// 返回脱敏后的邮箱地址
	return desensitizedEmail, nil
}

// DesensitizePhoneNumber 用于脱敏电话号码
func DesensitizePhoneNumber(phoneNumber string) (string, error) {
	// 使用正则表达式匹配电话号码
	re := regexp.MustCompile(`(\d{3})(-)(\d{3})(-)(\d{4})`)
	matches := re.FindStringSubmatch(phoneNumber)

	if len(matches) != 6 {
		return "", errors.New(consts.ErrPhoneFormat)
	}

	// 脱敏中间的部分，只显示第一个和最后一个数字，中间用星号(*)代替
	desensitizedPhoneNumber := matches[1] + "-" + "*****" + "-" + matches[5]

	// 返回脱敏后的电话号码
	return desensitizedPhoneNumber, nil
}
