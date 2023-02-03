package middleware

import (
	"net/http"
	"personal-collections/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Response(c *gin.Context) {
	c.Next()
	if c.Writer.Status() != 200 {
		return
	}
	obj, has := c.Get(resp.CTX_RESPONSE)
	if !has {
		logrus.Infof("[Response middleware] CTX_RESPONSE Not Found")
		// c.Status(http.StatusInternalServerError)
		return
	}
	resp, ok := obj.(resp.JsonResp)
	if !ok {
		c.Status(http.StatusInternalServerError)
		logrus.Errorf("[Response middleware] CTX_RESPONSE Type Error")
		return
	}

	if resp.ErrCode != 0 {
		logrus.Infof("Err serve '%s': [%d] %s", c.Request.URL.Path, resp.ErrCode, resp.ErrMsg)
	}

	c.JSON(http.StatusOK, resp)

}
