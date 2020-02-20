package middleware

import (
	"gin.weibo/pkg/flash"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OldValue 存储表单提交时数据的中间件
func OldValue() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			req := c.Request

			if req.Form == nil {
				req.ParseForm()
			}

			olaValue := make(map[string]string)
			for k, v := range c.Request.PostForm {
				olaValue[k] = v[0]
			}
			flash.SaveOldFormValue(c, olaValue)
		}

		c.Next()
	}
}
