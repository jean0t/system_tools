package main

import (
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
)

func getStat() string {
	var filePath string = "/proc/stat"
	var file, err = os.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(file)
}

// returns totalTime and idleTime, in this order
func getCpuTime() (uint64, uint64) {
	var totalTime, idleTime uint64 // same order of the returned variables
	var content = getStat()

	var timing = strings.Fields(strings.SplitN(content, "\n", 2)[0])[1:]
	for i, value := range timing {
		var number, _ = strconv.ParseUint(value, 10, 64)
		totalTime += number

		if i == 3 {
			idleTime = number
		}
	}

	return totalTime, idleTime
}

func calculateCpuUsage(interval time.Duration) (float64, error) {
	var totalTime1, idleTime1 uint64 = getCpuTime()

	time.Sleep(interval) // wait for the next state

	var totalTime2, idleTime2 uint64 = getCpuTime()

	var deltaIdle float64 = float64(idleTime2 - idleTime1)
	var deltaTotal float64 = float64(totalTime2 - totalTime1)
	if deltaTotal == 0 {
		return 0.0, nil
	}

	var usage float64 = 100 * (1.0 - deltaIdle/deltaTotal)

	return usage, nil

}

func main() {
	var cpuUsage, err = calculateCpuUsage(1 * time.Second)
	if err != nil {}

	fmt.Printf("%.1f%%", cpuUsage)
}
