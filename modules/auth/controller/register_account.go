package controller

import (
	"context"
	"firebase.google.com/go/auth"
	"golang-api-demo/initFirebase"
	"golang-api-demo/modules/auth/initModels"
	model "golang-api-demo/modules/auth/struct"
	"log"
)

func Register(register *model.Register) (*auth.UserRecord, error) {
	ctx := context.Background()

	client, _ := initFirebase.InitAuth()

	document := initModel.Register(register)

	id := document["id"].(string)
	email := document["email"].(string)
	password := document["password"].(string)

	params := (&auth.UserToCreate{}).
		UID(id).
		Email(email).
		Password(password)

	user, err := client.CreateUser(ctx, params)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
