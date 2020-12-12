package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FirestoreClient は、
type FirestoreClient struct {
	client *firestore.Client
	ctx    context.Context
}

// NewFirestoreClient は、
func NewFirestoreClient() *FirestoreClient {

	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FirestoreClient{
		client: client,
		ctx:    ctx,
	}
}
