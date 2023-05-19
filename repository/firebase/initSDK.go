package initFirebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Use a service account
func InitSDK() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("config/firebase_key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	//get all subscriptions
	iter := client.Collection("subscription").Documents(ctx)
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
	sub := client.Collection("subscription").Where("id", "==", "csBgld5GN7hn4lB0ISaX").Documents(ctx)
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
	/* _, _, err := client.Collection("subscription").Add(ctx, map[string]interface{}{
	  "email": "Jack@McMc.cz",
	  "id": "239kjl21kj312ljk232",
	  "newsletter_id": "13",
	})
	if err != nil {
	  log.Fatalf("Failed adding alovelace: %v", err)
	} */
}
