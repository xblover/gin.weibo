package home

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", map[string]interface{}{
		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}

// index2
func Index2(c *gin.Context) {
	c.HTML(http.StatusOK, "home2/index.html", map[string]interface{}{
		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}
