package test

import (
	"CloudStorage/core/define"
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	u, _ := url.Parse(define.BucketPath)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	key := "PreviewImage.png"

	_, _, err := client.Object.Upload(
		context.Background(), key, "PreviewImage.png", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileReader(t *testing.T) {
	u, _ := url.Parse(define.BucketPath)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretId,
			SecretKey: define.SecretKey,
		},
	})

	key := "PreviewImage.png"
	file, err := os.ReadFile("PreviewImage.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(file), nil,
	)
	if err != nil {
		panic(err)
	}
}
