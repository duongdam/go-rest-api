package userControllers

import (
	"github.com/gin-gonic/gin"
	"go-api/models/userModels"
	"net/http"
)

func GetAllUser(c *gin.Context) {
	users := userModels.GetAllUser()
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
