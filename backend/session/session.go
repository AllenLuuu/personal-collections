package session

import (
	"net/http"
	"personal-collections/config"
	"personal-collections/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	KEY_SESSION = "COLLECTIONS_SESSION"
	CTX_SESSION = "session"
)

func SessionMiddleware(c *gin.Context) {
	// NOTE: 必须是SameSite=None才能跨域Cookie！
	c.SetSameSite(http.SameSiteNoneMode)

	session_key, err := c.Cookie(KEY_SESSION)
	var session model.SessionData
	if err == nil {
		err = session.GetById(session_key)
	}

	// 如果找不到就新建
	if err != nil {
		if session.Create() != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		session_key = session.Id.Hex()
		logrus.Infof("Created Cookie Session: %s", session_key)
		c.SetCookie(KEY_SESSION, session_key, -1, "/", config.Server.Domain, true, true)
	} else {
		// logrus.Infof("Loaded Cookie Session: %s", session.Id)
		if session.User != nil {
			logrus.Infof("SessionUser: name=%s, id=%s", session.User.Username, session.User.Id)
		}
	}
	c.Set(CTX_SESSION, session)
	c.Next()
}

// panics if error
func GetSession(c *gin.Context) model.SessionData {
	data, ok := c.Get(CTX_SESSION)
	if ok {
		return data.(model.SessionData)
	} else {
		logrus.Fatalf("CTX_SESSION missing in gin.Context!")
	}
	return model.SessionData{}
}

func RemoveSession(c *gin.Context) {
	session_key, _ := c.Cookie(KEY_SESSION)
	model.SessionData{}.DeleteById(session_key)
	c.SetCookie(KEY_SESSION, "", 0, "/", config.Server.Domain, true, true)
}
