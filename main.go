package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers/userControllers"
)

func main() {
	router := gin.Default()
	router.GET("/all", userControllers.GetAllUser)

	router.Run("localhost:8080")
}
