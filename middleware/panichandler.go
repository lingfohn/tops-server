package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			requestId, _ := c.Get("RequestId")
			if err := recover(); err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message":   "内部服务错误",
					"requestId": requestId,
					"code":      500,
					"error":     err,
					"timestamp": time.Now().Unix(),
					"data":      map[string]string{},
				})
			}
		}()
		c.Next()
	}
}
