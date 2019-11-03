package result

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func GetPerformanceResult(logger *log.Logger, timeChannel chan float64) (max, min, avg, avg95 float64) {
	// 将channel 里的时间放入到slice中
	timeSlice := []float64{}
	for Time := range timeChannel {
		timeSlice = append(timeSlice, Time)
	}
	// sort slice
	sort.Float64s(timeSlice)
	fmt.Println(timeSlice)
	// get max
	MAX := getMax(timeSlice, logger)
	// get min
	MIN := getMin(timeSlice, logger)
	// get avg
	AVG := getAvg(timeSlice, logger)
	// get avg of 95%
	AVG95 := getAvg95(timeSlice, logger)

	return MAX, MIN, AVG, AVG95
}

// get max
func getMax(slice []float64, logger *log.Logger) float64 {
	sort.Float64s(slice)
	return slice[len(slice)-1]
}

// get min
func getMin(slice []float64, logger *log.Logger) float64 {
	sort.Float64s(slice)
	return slice[0]
}

// get ave
func getAvg(slice []float64, logger *log.Logger) float64 {
	sum := float64(0.00)
	for _,num := range slice {
		sum += num
	}
	return sum/float64(len(slice))
}

// get ave of 95%

func getAvg95(slice []float64, logger *log.Logger) float64 {
	lenth := len(slice) * 95 / 100
	sum := float64(0.00)
	for i:=0;i<lenth;i++ {
		sum += slice[i]
	}
	lenthFloat64, err := strconv.ParseFloat(strconv.Itoa(lenth), 64)
	if err != nil {
		logger.Println("transfer is bad", err)
		return -0.00
	}
	return sum/lenthFloat64
}