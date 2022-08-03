package initFirebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
	"path/filepath"
)

func InitAuth() *auth.Client {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("firebase-admin-sdk.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	conf := &firebase.Config{ProjectID: "brother-mori"}

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		panic("Firebase load error")
	}

	// Firebase Auth
	authenticate, err := app.Auth(ctx)
	if err != nil {
		panic("Firebase load error")
	}

	return authenticate
}

func InitFirestore() *firestore.Client {
	//Firebase admin SDK initialization
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("firebase-admin-sdk.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	//Firebase admin SDK initialization
	conf := &firebase.Config{ProjectID: "brother-mori"}

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		panic("Firebase load error")
	}

	// Firebase Firestore
	db, errFirestore := app.Firestore(ctx)

	if errFirestore != nil {
		log.Fatalln("Firebase load error", errFirestore)
	}

	return db
}
