package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nahidhasan98/kgc-crud/app/controller"
)

func CreateRoute(router *gin.Engine) {
	router.GET("/", controller.Index)

	authRG := router.Group("/auth")
	authRG.POST("/login", controller.Login)
	authRG.GET("/logout", controller.Logout)

	studentRG := router.Group("/api")
	studentRG.GET("/student", controller.GetAllStudents)
	studentRG.GET("/student/:id", controller.GetSingleStudent)
	studentRG.POST("/student", controller.CreateStudent)
	studentRG.PUT("/student/:id", controller.UpdateStudent)
	studentRG.DELETE("/student/:id", controller.DeleteStudent)
}
