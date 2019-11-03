package ceph

import (
	"github.com/minio/minio-go"
	"log"
)

func getAllObject(cephClient *minio.Client, logger *log.Logger) (objectSlice []string) {
	Slice := []string{}
	doneCh := make(chan struct{})
	defer close(doneCh)
	objectCh := cephClient.ListObjects(cephServer.bucket, "",cephServer.isRecursive,doneCh)
	for object := range objectCh {
		if object.Err != nil {
			logger.Println("get all object is is bad", object.Err)
			return nil
		}
		Slice = append(Slice, object.Key)
	}
	return Slice
}