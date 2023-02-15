package utils

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// TencentCos Tencent COS

// Constants
const (
	// TencentCosURL Tencent COS URL
	TencentCosURL = "https://tiktok-micro-1305132024.cos.ap-chengdu.myqcloud.com"
)

// TencentCos Tencent COS
type TencentCos struct {
	*cos.Client
}

// NewTencentCos New TencentCos
func NewTencentCos() *TencentCos {
	u, _ := url.Parse(TencentCosURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("TENCENT_SECRETID"),
			SecretKey: os.Getenv("TENCENT_SECRETKEY"),
		},
	})
	return &TencentCos{c}
}

// UploadFile Upload file to Tencent COS by stream
func (t *TencentCos) UploadFile(ctx context.Context, key string, body io.Reader) error {
	_, err := t.Object.Put(ctx, key, body, nil)
	return err
}
