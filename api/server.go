package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	auth "golang-api-demo/modules/auth/transport"
	groupController "golang-api-demo/modules/groups/controller"
	mysqlController "golang-api-demo/modules/mySQLDemo/controller"
	"golang-api-demo/util"
	"net/http"
	"os"
)

// Server serves HTTP requests for our banking service.

type Server struct {
	config util.Config
	router *gin.Engine
}

func InitServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.LoadHTMLFiles("static/index.html")

	if os.Getenv("GO_ENV") == "dev" {
		router.SetTrustedProxies([]string{"0.0.0.0"})
	}

	// get route /
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := router.Group("/v1")
	{

		{
			v1.POST("/register", auth.HandRegister)
			v1.POST("/login", auth.HandLogin)
		}

		// group route controller

		groupRouters := v1.Group("/groups").Use(authMiddleware())
		{
			groupRouters.POST("/create", groupController.HandleCreateGroup)
			groupRouters.GET("/list", groupController.HandleGetListGroup)
			groupRouters.GET("/getById", groupController.HandleGetOneGroup)
			groupRouters.POST("/update", groupController.HandleUpdateGroup)
			groupRouters.POST("/delete", groupController.HandleDeleteGroup)
		}

		mysqlRouters := v1.Group("/mySQLDemo").Use(authMiddleware())
		{
			mysqlRouters.GET("/", mysqlController.HandleGetData)
		}
	}

	server.router = router
	server.Start("0.0.0.0:" + os.Getenv("GO_PORT"))
}
