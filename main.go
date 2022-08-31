package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"golang-api-demo/api"
	"os"
)

func init() {
	err := gotenv.Load(".env")
	if err != nil {
		return
	}
	if os.Getenv("GO_ENV") != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	println("Running on port:", os.Getenv("GO_PORT"))
	// handle error when start server
	_, err := api.InitServer()
	if err != nil {
		fmt.Print("Init server error")
	}
}
