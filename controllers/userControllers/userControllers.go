package usersControllers

import (
	"github.com/gin-gonic/gin"
	"go-api/models/userModels"
	"net/http"
)

func GetAllUser(c *gin.Context) {
	users := userModels.GetAllUser()
	c.IndentedJSON(http.StatusOK, users)
}
