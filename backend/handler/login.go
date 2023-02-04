package handler

import (
	"personal-collections/model"
	"personal-collections/resp"
	"personal-collections/session"

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

	ss := session.GetSession(c)
	ss.User = &user
	ss.HasLoggedIn = true

	if err := ss.Update(); err != nil {
		logrus.Errorf("[handler] Failed to Save Session")
		resp.ERR(c, resp.E_DB_INSERT_ERROR, "cannot save session")
		return
	}

	logrus.Infof("[handler] User %s logged in, Session ID: %s", user.Username, ss.Id.Hex())
	resp.JSON(c, user)
}
