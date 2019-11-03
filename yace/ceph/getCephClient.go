package ceph

import (
	"github.com/minio/minio-go"
	"log"
)

// get ceph client
func getCephClient(logger *log.Logger) *minio.Client {
	client, err := minio.New(cephServer.endPoint, cephServer.accessKey, cephServer.secretKey, cephServer.useSSL)
	if err != nil {
		logger.Println("connect to ceph server is bad", err)
		return nil
	}
	return client
}
