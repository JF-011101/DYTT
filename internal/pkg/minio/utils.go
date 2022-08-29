/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-10 14:03:26
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:25:02
 * @FilePath: \dytt\pkg\minio\utils.go
 * @Description: Minio object stores business logic
 */

package minio

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/minio/minio-go/v7"
)

// CreateBucket 创建桶
func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		ilog.Error("bucketName invalid")
	}
	// location: https://cloud.google.com/storage/docs/locations
	location := "asia-east2"
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			ilog.Debugf("bucket %s already exists", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		ilog.Info("bucket create successfully")
	}
	return nil
}

// UploadLocalFile 上传本地文件（提供文件路径）至 minio
func UploadLocalFile(bucketName string, objectName string, filePath string, contentType string) (int64, error) {
	ctx := context.Background()
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		ilog.Errorf("localfile upload failed, %s", err)
		return 0, err
	}
	ilog.Infof("upload %s of size %d successfully", objectName, info.Size)
	return info.Size, nil
}

// UploadFile 上传文件（提供reader）至 minio
func UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	ctx := context.Background()
	n, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		ilog.Errorf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	ilog.Infof("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileUrl 从 minio 获取文件Url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		ilog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO: url可能要做截取
	return presignedUrl, nil
}
