package store

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

type Store struct {
	client *firestore.Client
}

// Use a service account
func NewConnection(ctx context.Context) *Store {
	sa := option.WithCredentialsFile("config/firebase_key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// defer client.Close()

	fb := &Store{
		client: client,
	}
	return fb
}

// Its like a class method.
func (fb *Store) GetSubscriptions(ctx context.Context) {
	//get all subscriptions
	iter := fb.client.Collection("subscription").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	//get subscription
	sub := fb.client.Collection("subscription").Where("id", "==", "csBgld5GN7hn4lB0ISaX").Documents(ctx)
	for {
		doc, err := sub.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	//new subscription TO FIX
	/* _, _, err := fb.client.Collection("subscription").Add(ctx, map[string]interface{}{
	     "email": "Jack@McMc.cz",
	     "id": "239kjl21kj312ljk232",
	     "newsletter_id": "13",
	   })
	   if err != nil {
	     log.Fatalf("Failed adding alovelace: %v", err)
	   } */
}
