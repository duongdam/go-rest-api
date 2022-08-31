package groupModel

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Group struct {
	Id             string      `json:"id" validate:"required"`
	ShortName      string      `json:"shortName" validate:"required"`
	Name           string      `json:"name" validate:"required"`
	Icon           string      `json:"icon" validate:"required"`
	LandId         string      `json:"landId" validate:"required"`
	SkillSender    SkillSender `json:"skillSender"`
	Owner          string      `json:"owner"`
	MemberIds      []string    `json:"memberIds"`
	OrganizationId string      `json:"organizationId"`
	CreatedAt      string      `json:"createdAt"`
	UpdatedAt      string      `json:"updatedAt"`
	IsDeleted      bool        `json:"isDeleted"`
}

type GroupId struct {
	Id string `uri:"id" binding:"required"`
}

type GroupIdJSON struct {
	Id string `json:"id" validate:"required"`
}

type SkillSender struct {
	Manager int `json:"manager"`
	Member  int `json:"members"`
}

type MemberGroup struct {
	UserId    string   `json:"user_id" validate:"required"`
	Avatar    string   `json:"avatar"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      string   `json:"role"`
	Tags      []string `json:"tags"`
	Skill     []string `json:"skill"`
	Sumori    []string `json:"sumori"`
	Komori    []string `json:"komori"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	IsDeleted bool     `json:"isDeleted"`
}

func GroupValidator(c *gin.Context) error {
	var group Group
	if err := c.ShouldBindJSON(&group); err == nil {
		validate := validator.New()
		if err := validate.Struct(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return err
		}
	}
	c.Next()
	return nil
}

func GroupIdValidator(c *gin.Context) error {
	var id GroupId
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"Error": err})
		c.Abort()
		return err
	}
	c.Next()
	return nil
}

func MemberValidator(c *gin.Context) error {
	var member MemberGroup
	if err := c.ShouldBindJSON(&member); err == nil {
		validate := validator.New()
		if err := validate.Struct(&member); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return err
		}
	}
	c.Next()
	return nil
}

func GroupIdJsonValidator(c *gin.Context) error {
	group := GroupIdJSON{}
	if err := c.BindJSON(&group); err == nil {
		validate := validator.New()
		if err := validate.Struct(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Check error": err.Error(),
			})
			c.Abort()
			return err
		}
	}
	c.Next()
	return nil
}
