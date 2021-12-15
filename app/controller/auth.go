package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/kgc-crud/app/model"
)

func Login(c *gin.Context) {
	// checking alreasy logged in or not
	session := sessions.Default(c)
	sessionUser := session.Get("user")
	if sessionUser != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "success",
			Message: "already logged in",
		})
		return
	}

	var user model.User
	err := c.ShouldBindJSON(&user)
	fmt.Println(user, err)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: "invalid JSON object",
		})
		return
	}
	fmt.Println("here")

	dbUser, err := model.Authenticate(&user)
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// saving session data
	session.Set("user", dbUser.ID)
	session.Save()

	c.JSON(http.StatusOK, &model.Response{
		Status:  "success",
		Message: "login successful",
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusOK, &model.Response{
			Status:  "error",
			Message: "logout unsuccessful",
		})
		return
	}
	// c.JSON(http.StatusOK, &model.Response{
	// 	Status:  "success",
	// 	Message: "logout successful",
	// })
	c.Redirect(http.StatusFound, "/")
}
