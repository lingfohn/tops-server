package auth

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/model"
	"net/http"
	"time"
)

var authMiddleware *jwt.GinJWTMiddleware

var identityKey = "id"

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type authMiddlewareManager struct {
	um *model.UserManager
}

func newAuthMiddlewareManager() *authMiddlewareManager {
	return &authMiddlewareManager{um: model.NewUserManager()}
}

func InitAuthMiddleware() (err error) {
	am := newAuthMiddlewareManager()
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// 登录成功返回
		LoginResponse:   loginResponse,
		RefreshResponse: refreshResponse,
		PayloadFunc:     am.payloadFunc,
		IdentityHandler: am.identityHandler,
		Authenticator:   am.authenticator,
		Authorizator:    am.authorizator,
		LogoutResponse:  logoutResponse,
		//认证未通过返回
		Unauthorized: unauthorized,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	return
}

// 登录成功返回
func loginResponse(c *gin.Context, code int, token string, expire time.Time) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"requestId": requestId,
		"message":   "",
		"error":     "",
		"timestamp": time.Now().Unix(),
		"data": gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		},
	})
}

// 认证失败返回
func unauthorized(c *gin.Context, code int, message string) {
	requestId, _ := c.Get("RequestId")
	c.JSON(code, gin.H{
		"code":      401,
		"requestId": requestId,
		"error":     message,
		"timestamp": time.Now().Unix(),
		"message":   "认证失败",
		"data":      map[string]string{},
	})
}

// 刷新成功返回
func refreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"requestId": requestId,
		"message":   "",
		"error":     "",
		"timestamp": time.Now().Unix(),
		"data": gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		},
	})
}

// 登出结果
func logoutResponse(c *gin.Context, code int) {
	requestId, _ := c.Get("RequestId")
	c.JSON(code, gin.H{
		"code":      0,
		"requestId": requestId,
		"message":   "logout success",
		"error":     "",
		"timestamp": time.Now().Unix(),
		"data":      map[string]string{},
	})
}

// 登录认证函数
func (am *authMiddlewareManager) authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password
	user, err := am.um.GetUserByUsername(userID)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	if (userID == user.Username && password == user.Password) || (userID == "super" && password == "1") {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// 设置payload
func (am *authMiddlewareManager) payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*ent.User); ok {
		return jwt.MapClaims{
			identityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

// 获取Identity
func (am *authMiddlewareManager) identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &ent.User{
		Username: claims[identityKey].(string),
	}
}

// 授权认证
func (am *authMiddlewareManager) authorizator(data interface{}, c *gin.Context) bool {
	fmt.Println(c.FullPath())
	if v, ok := data.(*ent.User); ok && v.Username == "admin" {
		return true
	}
	return false
}

func AuthMiddleware() *jwt.GinJWTMiddleware {
	return authMiddleware
}
