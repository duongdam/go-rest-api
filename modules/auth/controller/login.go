package controller

import (
	"context"
	"firebase.google.com/go/auth"
	"golang-api-demo/initFirebase"
	"golang-api-demo/modules/auth/initModels"
	model "golang-api-demo/modules/auth/struct"
	"log"
)

func Login(register *model.Register) (*auth.UserRecord, error) {

	ctx := context.Background()

	client, _ := initFirebase.InitAuth()

	document := initModel.Register(register)

	email := document["email"].(string)

	user, err := client.GetUserByEmail(ctx, email)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
