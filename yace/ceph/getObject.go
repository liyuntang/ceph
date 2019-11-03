package ceph

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func getObject(cephClient *minio.Client, logger *log.Logger, dataDir string, threadId int64, timeChannel chan float64)  {
	// create data file based threadId to download object

	dataFile := dataDir+"/"+fmt.Sprintf("%d", threadId)+".file"
	file, err := os.OpenFile(dataFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Println("create dataFile", dataFile, "is bad", err)
		return
	}
	// get object slice
	objectSlice := getAllObject(cephClient, logger)
	objectSliceLenth := len(objectSlice)
	// 根据 objectInfo.objectCount 设置的次数来get object
	var i int64
	for i=1;i<=objectInfo.objectCount;i++ {
		// 从objectSlice 中随机获取objectName
		index := rand.Intn(objectSliceLenth)
		objectName := objectSlice[index]
		bucketName := cephServer.bucket
		startTime := time.Now()
		objectValue, err1 := cephClient.GetObject(bucketName, objectName, minio.GetObjectOptions{})
		if err1 != nil {
			logger.Println("get object",objectName, "is bad", err1)
			return
		}
		// 将获取的object内容写入文件
		_, err2 := io.Copy(file, objectValue)
		if err2 != nil {
			logger.Println(objectName, "write to dataFile is bad", err2)
			return
		}
		runTime := time.Since(startTime)
		timeFloat64 := transfer(runTime.Nanoseconds()/1e6, logger)
		timeChannel <- timeFloat64
		logger.Println("hello my threadId is", threadId, "there are", objectInfo.objectCount, "need to run, and now run",i, "get object", objectName, "run time is",runTime)
	}
	// close dataFile
	file.Close()
}


