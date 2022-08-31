package groupController

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"
	"golang-api-demo/initFirebase"
	groupModel "golang-api-demo/modules/groups/model"
	"log"
	"net/http"
	"os"
	"time"
)

func HandleDeleteGroup(c *gin.Context) {
	group := groupModel.GroupIdJSON{}
	c.BindJSON(&group)
	response, err := DeleteGroup(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Delete error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteGroup(group *groupModel.GroupIdJSON) (*firestore.WriteResult, error) {
	//init ctx
	ctx := context.Background()
	// init client firestore when use
	client, _ := initFirebase.InitFirestore()
	// init batch write command
	groupRef := client.Collection(os.Getenv("GO_COL_GROUPS")).Doc(group.Id)
	// set field isDeleted true to remove
	response, err := groupRef.Set(ctx, map[string]interface{}{
		"isDeleted": true,
		"updatedAt": time.Now(),
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("Delete group error: %s", err)
		return nil, err
	}
	// close connection if not use anymore
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	return response, nil
}
