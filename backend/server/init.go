package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(r *gin.Engine) error {
	logrus.Info("[server] Initializing...")
	configRoutes(r)
	logrus.Info("[server] Initialized Successfully")
	return nil
}
