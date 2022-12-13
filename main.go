package main

import (
	"encoding/gob"
	"kafekoding-client/handlers"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(handlers.UserPayload{})
}

func main() {
	r := gin.Default()
	r.HTMLRender = createMyRender()

	r.Use(sessions.Sessions("kafekoding-client-id", cookie.NewStore([]byte("mysecretkey"))))

	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)

	r.GET("/", handlers.Index)
	r.GET("/test", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")
		ctx.JSON(http.StatusOK, user)
	})
	r.GET("/classes/:id", handlers.Detail)

	_ = r.Run(":3000")
}
