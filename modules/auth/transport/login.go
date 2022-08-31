package transport

import (
	"github.com/gin-gonic/gin"
	authController "golang-api-demo/modules/auth/controller"
	model "golang-api-demo/modules/auth/struct"
	"net/http"
)

func HandLogin(c *gin.Context) {
	var register model.Register
	if err := c.ShouldBind(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := authController.Login(&register)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"created user data": response})
}
