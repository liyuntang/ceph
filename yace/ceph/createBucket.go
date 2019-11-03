package ceph

import (
	"github.com/minio/minio-go"
	"log"
)

func createBucket(cephClient *minio.Client, logger *log.Logger)  {
	bucketName := cephServer.bucket
	err := cephClient.MakeBucket(bucketName, cephServer.location)
	if err != nil {
		logger.Println("create bucket", bucketName, "is bad", err)
		return
	}
	logger.Println("create bucket", bucketName, "is ok")
}
