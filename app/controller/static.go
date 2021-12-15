package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	loggedIn := false

	session := sessions.Default(c)
	sessionID := session.Get("user")
	if sessionID != nil {
		loggedIn = true
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"LoggedIn": loggedIn,
		"Title":    "Dept. of Management | KGC",
	})
}
