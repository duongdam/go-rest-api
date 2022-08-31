package initFirebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

func FirebaseInit() (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "[ProjectID]"}
	opt := option.WithCredentialsFile("initFirebase/firebase2-admin-sdk.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	}
	return app, nil
}

func InitAuth() (*auth.Client, error) {
	app, _ := FirebaseInit()
	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return nil, errors.Wrap(errAuth, "error initializing firebase auth (creating client)")
	}

	return client, nil
}

func InitFirestore() (*firestore.Client, error) {
	app, err := FirebaseInit()
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firestore (creating client)")
	}
	client, err := app.Firestore(context.Background())
	return client, nil
}
