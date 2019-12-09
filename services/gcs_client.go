package services

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/Nks/go-gcs-ffmpeg/models"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Client Main struct
type GcsClient struct {
	context context.Context
	client  storage.Client
}

//Create and retrieve google cloud storage client
func (gcs *GcsClient) CreateClient() error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	if err != nil {
		return fmt.Errorf("unable create client %s", err)
	}

	gcs.setContext(ctx)
	gcs.setClient(*client)

	return nil
}

func (gcs *GcsClient) CreateSelfSignedUrlForFile(bucket string, filename string, serviceAccount string) (string, error) {
	jsonKey, err := ioutil.ReadFile(serviceAccount)
	if err != nil {
		return "", fmt.Errorf("cannot read the JSON key file, err: %v", err)
	}

	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(24 * time.Hour),
	}

	u, err := storage.SignedURL(bucket, filename, opts)
	if err != nil {
		return "", fmt.Errorf("unable to generate a signed URL: %v", err)
	}

	return u, nil
}

func (gcs *GcsClient) UploadStreamToGcs(tempPath string, parameters *models.Parameters, public bool) error {
	bucket := parameters.Storage
	var files []string

	err := filepath.Walk(tempPath, scanFolder(&files))

	if err != nil {
		return fmt.Errorf("unable list stream path with error %v", err)
	}

	done := make(chan struct{})
	count := uint64(0)

	for _, file := range files {
		count += 1

		go func(file string) {
			defer func() { done <- struct{}{} }()

			var r io.Reader
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			r = f

			name := strings.Replace(file, tempPath, parameters.Output, -1)

			ctx := context.Background()
			_, err = upload(ctx, r, bucket, name, public)

			if err != nil {
				switch err {
				case storage.ErrBucketNotExist:
					log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
				default:
					log.Fatal(err)
				}
			} else {
				log.Println("File uploaded:", name)
			}

		}(file)
	}

	for i := uint64(0); i < count; i++ {
		<-done
	}

	return nil
}

func (gcs *GcsClient) GetClient() storage.Client {
	return gcs.client
}

func scanFolder(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

func upload(ctx context.Context, r io.Reader, bucket, name string, public bool) (*storage.ObjectHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bh := client.Bucket(bucket)

	if _, err = bh.Attrs(ctx); err != nil {
		return nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, err
		}
	}

	extension := filepath.Ext(name)
	var contentType string

	switch extension {
	case ".m3u8", ".m3u":
		contentType = "application/vnd.apple.mpegurl"
	case ".ts", ".tsa", ".tsv":
		contentType = "video/mp2t"
	default:
		contentType = ""
	}

	if contentType != "" {
		objectAttrsToUpdate := storage.ObjectAttrsToUpdate{
			ContentType:  contentType,
			CacheControl: "private, max-age=21296", //youtube's cache control
		}

		if _, err := obj.Update(ctx, objectAttrsToUpdate); err != nil {
			return obj, err
		}
	}

	return obj, err
}

func (gcs *GcsClient) setContext(ctx context.Context) {
	gcs.context = ctx
}

func (gcs *GcsClient) setClient(v storage.Client) {
	gcs.client = v
}
