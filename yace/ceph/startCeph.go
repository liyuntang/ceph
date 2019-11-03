package ceph

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"ceph/yace/common"
)

// ceph server configration
type ceph struct {
	endPoint string
	accessKey string
	secretKey string
	bucket string
	location string
	useSSL bool
	isRecursive bool
}

// object info
type object struct {
	objectSize int64
	objectOption string
	objectCount int64
	dataDir string
}

var (
	cephServer ceph
	objectInfo object
)



func StartCeph(config *common.TomlConfig, logger *log.Logger, threadId int64, group *sync.WaitGroup, timeChannel chan float64)  {
	// 初始化ceph参数及object信息
	cephConfig(config)
	option(config)
	// get datadir info
	dataDir := config.Yace.DataDir
	// get ceph client
	cephClient := getCephClient(logger)
	// get object option
	if strings.ToLower(objectInfo.objectOption) == "get" {
		// do get object test
		fmt.Println("start get object test")
		getObject(cephClient, logger, dataDir, threadId, timeChannel)
	} else if strings.ToLower(objectInfo.objectOption) == "put" {
		// do put object test
		fmt.Println("start put object test")
		putObject(cephClient, logger, threadId, timeChannel)
	} else if strings.ToLower(objectInfo.objectOption) == "create" {
		// do create bucket
		createBucket(cephClient, logger)
	} else {
		logger.Println("plese tell me you want: get object or put object")
	}
	group.Done()
}


// init ceph server
func cephConfig(config *common.TomlConfig) {
	cephServer.endPoint = config.Ceph_server.Address + ":" + fmt.Sprintf("%d", config.Ceph_server.Port)
	cephServer.accessKey = config.Ceph_server.AccessKey
	cephServer.secretKey = config.Ceph_server.SecretKey
	cephServer.bucket = config.Ceph_server.Bucket
	cephServer.useSSL = false
	cephServer.isRecursive = true
}

func option(config *common.TomlConfig)  {
	objectInfo.objectOption = config.Yace.Option
	objectInfo.objectSize = config.Yace.ObjectSize
	objectInfo.objectCount = config.Yace.ObjectCount
	objectInfo.dataDir = config.Yace.DataDir
}