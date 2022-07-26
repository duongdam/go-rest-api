package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers/userControllers"
)

func main() {
	router := gin.Default()
	router.GET("/api/user/list", userControllers.GetAllUser)
	router.GET("/api/user/get", userControllers.GetUser)
	router.POST("/api/user/create", userControllers.CreateUser)
	router.POST("/api/user/update", userControllers.UpdateUser)
	router.POST("/api/user/delete", userControllers.DeleteUser)

	router.Run("localhost:8080")
}
