package groupController

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang-api-demo/initFirebase"
	groupModel "golang-api-demo/modules/groups/model"
	"google.golang.org/api/iterator"
	"net/http"
	"os"
)

func HandleGetListGroup(c *gin.Context) {
	response, err := GetListGroup()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, response)
}

func HandleGetOneGroup(c *gin.Context) {
	if err := groupModel.GroupIdValidator(c); err != nil {
		return
	}
	id := c.Param("id")
	response, err := GetOneGroup(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetListGroup() ([]map[string]interface{}, error) {
	//init ctx
	ctx := context.Background()
	var bs []map[string]interface{}
	// init client firestore when use
	client, _ := initFirebase.InitFirestore()

	iter := client.Collection(os.Getenv("GO_COL_GROUPS")).Where("isDeleted", "==", false).Documents(ctx)
	defer iter.Stop() // add this line to ensure resources cleaned up
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var b map[string]interface{}
		if err := doc.DataTo(&b); err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}

	return bs, nil
}

func GetOneGroup(id string) (map[string]interface{}, error) {
	//init ctx
	ctx := context.Background()
	// init client firestore when use
	client, _ := initFirebase.InitFirestore()

	snapshot, err := client.Collection(os.Getenv("GO_COL_GROUPS")).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	m := snapshot.Data()
	if m["isDeleted"] == true {
		return nil, errors.New("group has deleted")
	}

	return m, nil
}
