package middleware

import (
	"personal-collections/resp"
	"personal-collections/session"

	"github.com/gin-gonic/gin"
)

// 检查医生和管理端登录状况
// 医生还要检查是否是审核通过的
func CheckUserLoggedIn(c *gin.Context) {
	ss := session.GetSession(c)
	if ss.User == nil || !ss.HasLoggedIn {
		resp.ERR(c, resp.E_NOT_LOGGED_IN, "not logined!")
		c.Abort()
		return
	}
	c.Next()
}
