package ceph

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"log"
	"os"
	"time"
)

func putObject(cephClient *minio.Client, logger *log.Logger, threadId int64, timeChannel chan float64)  {

	// get bucketName
	bucketName := cephServer.bucket
	// get objectSize
	objectSize := objectInfo.objectSize
	// 根据object 大小获取io.reader
	fileReader := getIoReader(objectSize, logger,threadId)
	// 压测次数
	runTimes := objectInfo.objectCount
	var i int64
	for i=1;i<=runTimes;i++ {
		// get objectName
		objectName := getObjectName(logger, threadId)
		// put object
		startTime := time.Now()
		_, err := cephClient.PutObject(bucketName, objectName, fileReader, objectSize, minio.PutObjectOptions{ContentType:"application/octet-stream"})
		if err != nil {
			logger.Println("put object", objectName, "to ceph is bad", err)
			return
		}
		runTime := time.Since(startTime)
		timeFloat64 := transfer(runTime.Nanoseconds()/1e6, logger)
		timeChannel <- timeFloat64
		logger.Println("hello my threadId is", threadId, "there are", objectInfo.objectCount, "need to run, and now run",i, "put object", objectName, "to ceph cluster is ok,run time is",runTime)
	}

}

func getObjectName(logger *log.Logger, threadId int64) string {
	objectName := fmt.Sprintf("%d-%d", threadId, time.Now().Nanosecond())
	return objectName
}

func getIoReader(objectSize int64, logger *log.Logger, threadId int64) io.Reader {
	dataFile := objectInfo.dataDir+"/"+fmt.Sprintf("%d", threadId)+".file"
	objectValue := make([]byte, objectSize)
	_, err := os.Stat(dataFile)
	if err == nil {
		// 说明文件存在，删除文件
		err1 := os.Remove(dataFile)
		if err1 != nil {
			logger.Println("dataFile",dataFile, "is exsit, please remove it",err1)
		}
	}
	file, err2 := os.OpenFile(dataFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		logger.Println("create dataFile is bad", err)
		return nil
	}
	// 写入数据
	_, err3 := file.Write(objectValue)

	if err3 != nil {
		logger.Println("write object to dataFile is bad", err3)
		return nil
	}
	// open file get io.reader
	f, err4 := os.Open(dataFile)
	if err4 != nil {
		logger.Println("open dataFile", dataFile, "is bad", err4)
	}
	return f
}