package store

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
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
func (fb *Store) GetSubscriptions(ctx context.Context) []Subscription {
	var subscriptions []Subscription
	iter := fb.client.Collection("subscription").Documents(ctx)

	for {
    doc, err := iter.Next()
    if err == iterator.Done {
        break
    }
    if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
    }
    var subscription Subscription
    if err := doc.DataTo(&subscription); err != nil {
			log.Fatalf("Failed to map subscription to struct: %v", err)
    }
    subscriptions = append(subscriptions, subscription)
	}
	return subscriptions
}

func (fb *Store) GetSubscription(ctx context.Context, newsletter_id int, email string) []Subscription {
	var subscriptions []Subscription

	iter := fb.client.Collection("subscription").Where("newsletter_id", "==", newsletter_id).Where("email", "==", email).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
				break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var subscription Subscription
		if err := doc.DataTo(&subscription); err != nil {
			log.Fatalf("Failed to map subscription to struct: %v", err)
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions
}

func (fb *Store) NewSubscription(ctx context.Context, newsletter_id int, email string) error{

	id := uuid.New().String()
	_, err := fb.client.Collection("subscription").Doc(id).Set(ctx, map[string]interface{}{
		 "email": email,
	   "id": id,
	   "newsletter_id": newsletter_id,
	})
	if err != nil {
	  log.Fatalf("Failed adding subscription: %v", err)
	} 

	return err
}

func (fb *Store) DeleteSubscription(ctx context.Context, id string) string{

	_, err := fb.client.Collection("subscription").Doc(id).Delete(ctx)
	if err != nil {
	  log.Fatalf("Failed deleting subscription: %v", err)
	}

	return "Successfully deleted"
}
