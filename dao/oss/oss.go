package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint        = "your-oss-endpoint"
	accessKeyID     = "your-access-key-id"
	accessKeySecret = "your-access-key-secret"
	bucketName      = "your-bucket-name"
	client          *oss.Client
)

func Init() (err error) {
	client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}
	return
}

// UploadAvatarToOSS 上传用户头像到阿里云 OSS
func UploadAvatarToOSS(imagePath, imageName string) (err error) {
	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 上传图片到 OSS
	err = bucket.PutObjectFromFile(imageName, imagePath)
	if err != nil {
		return err
	}

	return nil
}

// UploadCommentImageToOSS 上传用户评论的图片到阿里云 OSS
func UploadCommentImageToOSS(imagePath, imageName string) (err error) {
	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 上传图片到 OSS
	err = bucket.PutObjectFromFile(imageName, imagePath)
	if err != nil {
		return err
	}

	return nil
}
