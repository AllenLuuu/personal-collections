package server

import (
	. "personal-collections/handler"
	"personal-collections/middleware"
	"personal-collections/session"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func configRoutes(e *gin.Engine) {
	e.GET("/ping", ping)

	cors_cfg := cors.Config{
		AllowOrigins:     []string{"https://allenluuu.com", "http://127.0.0.1:5173", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           3600,
		ExposeHeaders:    []string{"Authorization", "Set-Cookie"},
	}
	e.Use(cors.New(cors_cfg))

	root := e.Group("/", session.SessionMiddleware, middleware.Response)
	root.POST("/login", Login)

	// 以下是需要登录的路由
	// admin := root.Group("/admin", middleware.CheckUserLoggedIn)
}
