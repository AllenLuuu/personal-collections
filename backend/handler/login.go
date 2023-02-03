package handler

import (
	"personal-collections/model"
	"personal-collections/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var req model.User
	var user model.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_BAD_PARAM, "参数错误")
		return
	}

	err = user.GetByUsername(req.Username)
	if err != nil {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_WRONG_USERPASS, "用户名或密码错误")
		return
	}

	if user.Password != req.Password {
		logrus.Info(err.Error())
		resp.ERR(c, resp.E_WRONG_USERPASS, "用户名或密码错误")
		return
	}

	resp.JSON(c, user)
}
