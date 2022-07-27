package userControllers

import (
	"github.com/gin-gonic/gin"
	"go-api/models/userModels"
	"log"
	"net/http"
	"os"
)

func GetAllUser(c *gin.Context) {
	users := userModels.GetAllUser()
	nodeEnv := os.Getenv("NODE_ENV")

	log.Println("nodeEnv", nodeEnv)

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserQuery(c *gin.Context) {
	id := c.Query("id")
	users := userModels.GetAllUser()
	nodeEnv := os.Getenv("NODE_ENV")

	log.Println("id", id)
	log.Println("users", users)
	log.Println("nodeEnv", nodeEnv)

	// Loop over the list of users, looking for user matched
	for _, a := range users {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func GetUserParam(c *gin.Context) {
	id := c.Param("id")
	users := userModels.GetAllUser()
	nodeEnv := os.Getenv("NODE_ENV")

	log.Println("id", id)
	log.Println("users", users)
	log.Println("nodeEnv", nodeEnv)

	// Loop over the list of users, looking for user matched

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
