package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/kgc-crud/app/router"
	"github.com/nahidhasan98/kgc-crud/config"
)

func cookieInit(r *gin.Engine) {
	store := cookie.NewStore([]byte(config.CookieSecretKey))
	r.Use(sessions.Sessions("kgc", store))
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	cookieInit(r)

	r.LoadHTMLGlob("view/*")
	r.Static("/assets", "./assets")

	router.CreateRoute(r)

	r.Run(":6001")
	fmt.Println("Server running on port 6001...")
}
