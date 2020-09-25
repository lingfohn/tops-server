package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/svcerr"
	"io"
	"net/http"
	"os"
	"time"
)

type JSONRenderBody struct {
	Message   string      `json:"message"`
	RequestId string      `json:"requestId"`
	Code      int         `json:"code"`
	Error     string      `json:"error"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	Query     interface{} `json:"query,omitempty"`
}

var (
	defaultMapValue           = map[string]string{}
	out             io.Writer = os.Stdout
)

func JSONRenderWrap(next func(ctx *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := next(c)
		requestId, ok := c.Get("RequestId")
		if err != nil {
			resp := JSONRenderBody{Data: defaultMapValue}
			switch err.(type) {
			case *svcerr.SvcError:
				se, _ := err.(*svcerr.SvcError)
				resp.Code = se.Code()
				resp.Message = se.Error()
				resp.Error = se.Unwrap().Error()
			default:
				resp.Code = svcerr.InternalServerError.Code()
				resp.Message = "发生未知错误"
				resp.Error = err.Error()
			}
			resp.RequestId = requestId.(string)
			resp.Timestamp = time.Now().Unix()
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		result, ok := c.Get("data")
		if !ok {
			result = defaultMapValue
		}
		message, ok := c.Get("message")
		if !ok {
			message = ""
		}
		query, ok := c.Get("query")
		if !ok {
			query = defaultMapValue
		}

		total, ok := c.Get("total")
		if !ok {
			total = 0
		}
		resp := JSONRenderBody{
			Message:   message.(string),
			RequestId: requestId.(string),
			Code:      0,
			Error:     "",
			Timestamp: time.Now().Unix(),
			Data:      result,
			Query:     query,
			Total:     total.(int),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
}
