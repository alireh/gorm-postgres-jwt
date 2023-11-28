package main

import (
	"fmt"
	"gorm-postgres-jwt/controller"
	"gorm-postgres-jwt/middleware"
	"gorm-postgres-jwt/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	model.SetDBClient()
}
func main() {
	fmt.Println("Welcome to Go authorization with Go")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Home router",
		})
	})

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/api/v1", middleware.Authorize, controller.Resources)

	r.Run(":5000")
}
