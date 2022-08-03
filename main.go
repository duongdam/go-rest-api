package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"go-api/controllers/userControllers"
	"go-api/initFirebase"
	"net/http"
	"os"
)

func init() {
	gotenv.Load(".env")
	if os.Getenv("GO_ENV") != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	println("Running on port:", os.Getenv("GO_PORT"))
	router := gin.Default()
	router.LoadHTMLFiles("static/index.html")

	if os.Getenv("GO_ENV") == "dev" {
		router.SetTrustedProxies([]string{"0.0.0.0"})
	}

	// create/configure database instance
	db := initFirebase.InitFirestore()
	auth := initFirebase.InitAuth()

	// set db to gin context with a middleware to all incoming request
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("auth", auth)
	})

	// get route /
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/api/user/list", userControllers.GetAllUser)
	router.GET("/api/user/get", userControllers.GetUserQuery)
	router.GET("/api/user/get/:id", userControllers.GetUserParam)
	router.POST("/api/user/create", userControllers.CreateUser)
	router.POST("/api/user/update", userControllers.UpdateUser)
	router.POST("/api/user/delete", userControllers.DeleteUser)

	router.Run("0.0.0.0:" + os.Getenv("GO_PORT"))
}
