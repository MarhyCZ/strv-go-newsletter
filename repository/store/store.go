package store

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
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
			log.Fatalln("Failed to map subscription to struct: %v", err)
    }
    subscriptions = append(subscriptions, subscription)
}
return subscriptions
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
				log.Fatalln("Failed to map subscription to struct: %v", err)
			}
			subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}
