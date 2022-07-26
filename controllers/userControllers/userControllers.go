package userControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/models/userModels"
	"net/http"
)

func GetAllUser(c *gin.Context) {
	users := userModels.GetAllUser()
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	users := userModels.GetAllUser()

	fmt.Println("id", id)
	fmt.Println("users", users)

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for _, a := range users {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func CreateUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
