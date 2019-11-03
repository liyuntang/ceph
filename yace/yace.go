package main

import (
	"fmt"
	"sync"
	"time"
	"ceph/yace/ceph"
	"ceph/yace/common"
	"ceph/yace/result"
)
var waitgroup sync.WaitGroup


func main()  {
	// get configration
	tomlConfig := common.Config("src/yace/conf/ceph.conf")
	// init log
	logger := common.InitLogger()

	// make a channel
	allTimes := tomlConfig.Yace.ObjectCount * tomlConfig.Yace.Parallel
	timeChannel := make(chan float64, allTimes)

	// 根据 parallel设置开启并发访问
	threads := tomlConfig.Yace.Parallel
	fmt.Println("performance test is start")
	startTime := time.Now()
	var threadId int64
	for threadId =1;threadId<= threads;threadId++ {
		waitgroup.Add(1)
		fmt.Println("start thread",threadId)
		go ceph.StartCeph(tomlConfig, logger, threadId, &waitgroup, timeChannel)
	}
	waitgroup.Wait()
	// process is over
	runTime := time.Since(startTime)
	fmt.Println("performance test is over, run time is", runTime)
	// close channel
	close(timeChannel)

	// process performance result
	max, min, avg, avg95 := result.GetPerformanceResult(logger, timeChannel)
	logger.Println("并发数为",threads,"请求总量为", allTimes ,"压测时间为",runTime, "object大小为",tomlConfig.Yace.ObjectSize/1000,"K")
	logger.Println("最大响应时间为",max, "ms,最小响应时间为", min, "ms,平均响应时间为", avg, "ms,95%的请求响应平均时间为", avg95, "ms")

}
