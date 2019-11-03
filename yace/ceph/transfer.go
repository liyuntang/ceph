package ceph

import (
	"log"
	"strconv"
)

func transfer(num int64, logger *log.Logger) float64 {
	numFloat64, err := strconv.ParseFloat(strconv.FormatInt(num, 10), 64)
	if err != nil {
		logger.Println("transfer num to float64 is bad", err)
		return -0.0
	}
	return numFloat64

}
