package main

import (
	"fmt"
	"os"
	"personal-collections/config"
	"personal-collections/database"
	"personal-collections/model"
	"personal-collections/server"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// err handler
func Do(err error) {
	if err != nil {
		logrus.Fatalf("Error in Init: %s", err)
	}
}

func initLogger() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
		ForceQuote:                true,
		EnvironmentOverrideColors: false,
	})
	return nil
}

func main() {
	initLogger()
	Do(config.Init())
	Do(database.Init())
	Do(model.Init())

	r := gin.Default()
	Do(server.Init(r))

	err := r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	if err != nil {
		logrus.Errorf("Error While Running Server: %s", err.Error())
		os.Exit(-1)
	}
}
