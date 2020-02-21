package flash

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ValidateContextAndCookieKeyName = "validatorMessage"
	ValidateSeparator               = "$$|$$"
)

// SaveValidateMessage 存储参数验证的错误信息
func SaveValidateMessage(c *gin.Context, msgArr []string) {
	f := NewFlashByName(ValidateContextAndCookieKeyName)
	f.Data = map[string]string{
		"errors": strings.Join(msgArr, ValidateSeparator),
	}
	f.save(c, ValidateContextAndCookieKeyName)
}

// ReadValidateMessage 读取参数验证的错误信息
func ReadValidateMessage(c *gin.Context) []string {
	errorStr := read(c, ValidateContextAndCookieKeyName).Data["errors"]

	if errorStr == "" {
		return []string{}
	}
	// 不做上面的判断， Split 切分空字符串会得[""]
	return strings.Split(errorStr, ValidateSeparator)
}
