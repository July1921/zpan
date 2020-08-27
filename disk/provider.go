package disk

import (
	"fmt"
	"time"
)

type Object struct {
	Key          string    `json:"key"`
	Dir          bool      `json:"dir"`
	Type         string    `json:"type"`
	Size         int64     `json:"size"`
	ETag         string    `json:"etag"`
	LastModified time.Time `json:"last_modified"`
}

type Objects []Object

type Provider interface {
	SetLifecycle(bucketName string) error
	BuildCallback(url, body string) string
	UploadURL(bucketName, filename, objectKey, contentType, callback string, publicRead bool) (url string, headers map[string]string, err error)
	DownloadURL(bucketName, objectKey string) (url string, err error)
	ListObject(bucketName, prefix, marker string, limit int) (objects Objects, nextMarker string, err error)
	TagRename(bucketName, objectKey, filename string) error
	TagDelObject(bucketName, objectKey string) error
	DeleteObject(bucketName, objectKey string) error
	DeleteObjects(bucketName string, objectKeys []string) error
}

type Config struct {
	Name         string
	Bucket       string
	Endpoint     string
	AccessKey    string
	AccessSecret string
}

type ProviderConstructor func(provider Config) (*AliOss, error)

var providerConstructors = map[string]ProviderConstructor{
	"ali-oss": newAliOss,
}

func New(provider Config) (Provider, error) {
	if providerConstructor, ok := providerConstructors[provider.Name]; ok {
		return providerConstructor(provider)
	}

	return nil, fmt.Errorf("provider %s not found", provider.Name)
}
