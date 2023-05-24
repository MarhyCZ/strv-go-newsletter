package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	//"encoding/base64"

	storageFn "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	storage "firebase.google.com/go/v4/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Storage struct {
	client *storage.Client
}

func NewConnection(ctx context.Context) *Storage {
	config := &firebase.Config{
		StorageBucket: "strv-go-newsletter.appspot.com",
	}
	opt := option.WithCredentialsFile("config/firebase_key.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	fb := &Storage{
		client: client,
	}

	return fb
	// 'bucket' is an object defined in the cloud.google.com/go/storage package.
	// See https://godoc.org/cloudgoogle.com/go/storage#BucketHandle
	// for more details.storage.go
}

func (fb *Storage) GetIssuesList(ctx context.Context, w io.Writer, delim string, prefix string) error {
	bucket, err := fb.client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	it := bucket.Objects(ctx, &storageFn.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%v).Objects: %w", bucket, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}

func (fb *Storage) DownloadFileIntoMemory(ctx context.Context, w io.Writer, issue string) ([]byte, error) {
	bucket, err := fb.client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	rc, err := bucket.Object(issue).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %w", issue, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	fmt.Fprintf(w, "Blob %v downloaded.\n", issue)
	return data, nil
}

func (fb *Storage) StreamFileUpload(ctx context.Context, w io.Writer, issue string, data []byte) error {
	bucket, err := fb.client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	//md, _ := base64.StdEncoding.DecodeString(data)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := bucket.Object(issue).NewWriter(ctx)
	wc.ContentType = "text/markdown"
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := wc.Write(data); err != nil {
		return fmt.Errorf("Writer.Write: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	fmt.Fprintf(w, "%v uploaded.\n", issue)

	return nil
}
