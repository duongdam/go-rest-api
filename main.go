package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers/userControllers"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/api/user/list", userControllers.GetAllUser)
	router.GET("/api/user/get", userControllers.GetUserQuery)
	router.GET("/api/user/get/:id", userControllers.GetUserParam)
	router.POST("/api/user/create", userControllers.CreateUser)
	router.POST("/api/user/update", userControllers.UpdateUser)
	router.POST("/api/user/delete", userControllers.DeleteUser)

	router.Run("localhost:8080")
}
