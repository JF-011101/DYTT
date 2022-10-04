/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-10 14:03:26
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-16 11:44:36
 * @FilePath: \dytt\pkg\minio\init.go
 * @Description: Minio object storage initialization
 */

package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
)

var (
	minioClient          *minio.Client
	Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
	MinioEndpoint        = Config.Viper.GetString("minio.Endpoint")
	MinioAccessKeyId     = Config.Viper.GetString("minio.AccessKeyId")
	MinioSecretAccessKey = Config.Viper.GetString("minio.SecretAccessKey")
	MinioUseSSL          = Config.Viper.GetBool("minio.UseSSL")
	MinioVideoBucketName = Config.Viper.GetString("minio.VideoBucketName")
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		ilog.Errorf("minio client init failed: %v", err)
	}
	ilog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		ilog.Errorf("minio client init failed: %v", err)
	}
}
