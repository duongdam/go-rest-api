package groupController

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"
	"golang-api-demo/initFirebase"
	initGroup "golang-api-demo/modules/groups/firestore"
	groupModel "golang-api-demo/modules/groups/model"
	"log"
	"net/http"
	"os"
)

func HandleUpdateGroup(c *gin.Context) {
	var group groupModel.Group
	if err := groupModel.GroupValidator(c); err != nil {
		return
	}

	response, err := UpdateGroup(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Update group data": response})
}

func UpdateGroup(group *groupModel.Group) (map[string]interface{}, error) {
	//init ctx
	ctx := context.Background()

	// init client firestore when use
	client, _ := initFirebase.InitFirestore()

	// init batch write command
	batch := client.Batch()

	// parse data by format of firebase document
	document, id := initGroup.ParseGroup(group)
	// set data to cloud firestore
	groupRef := client.Collection(os.Getenv("GO_COL_GROUPS")).Doc(id)
	batch.Set(groupRef, document, firestore.MergeAll)

	_, err := batch.Commit(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	// close connection if not use anymore
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	return document, nil
}
