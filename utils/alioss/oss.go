package alioss

import (
	"fmt"
	"io"
	"net/url"
	"strings"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// OSS .
type OSS struct {
	client *oss.Client
	bucket *oss.Bucket
}

var (
	syncOnce sync.Once
	_oss     *OSS
)

// Config .
type Config struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

// New .
func New(c *Config) *OSS {
	syncOnce.Do(func() {
		_oss = new(OSS)
		client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
		if err != nil {
			logx.Error(err)
			panic(err)
		}
		_oss.client = client
		bucket, err := client.Bucket(c.BucketName)
		if err != nil {
			logx.Error(err)
			panic(err)
		}
		_oss.bucket = bucket
	})
	return _oss
}

// Put 上传.
func (oss *OSS) Put(objectKey string, reader io.Reader, fileName string) (uri string, err error) {
	err = oss.bucket.PutObject(objectKey, reader)
	if err != nil {
		logx.Error(err)
	}
	// 文件名编码返回url
	objKey := strings.Replace(objectKey, fileName, url.PathEscape(fileName), 1)
	// uri = objKey
	uri = fmt.Sprintf("https://%s.%s/%s", oss.bucket.BucketName, strings.Replace(oss.client.Config.Endpoint, "https://", "", 1), objKey)
	fmt.Println(uri)
	return
}

// Delete 删除.
func (oss *OSS) Delete(objectKey string) (err error) {
	err = oss.bucket.DeleteObject(objectKey)
	if err != nil {
		logx.Error(err)
	}
	return
}
