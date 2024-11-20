package main

import (
	"fmt"
	"math"
	"golang.org/x/sys/unix"
)

//used ram is returned
func getUsedRam() uint64 {
	var ramInfo unix.Sysinfo_t

	if err := unix.Sysinfo(&ramInfo); err != nil {
		return 0
	}

	var totalRam uint64 = ramInfo.Totalram * uint64(ramInfo.Unit)
	var freeRam uint64 = ramInfo.Freeram * uint64(ramInfo.Unit)
	var sharedRam uint64 = ramInfo.Sharedram * uint64(ramInfo.Unit)
	var bufferRam uint64 = ramInfo.Bufferram * uint64(ramInfo.Unit)
	var usedRam uint64 = totalRam - (freeRam + bufferRam + sharedRam)

	return usedRam
}

func main() {
	var convertGB float64 = math.Pow(2, 30)
	var ramUsed = float64(getUsedRam()) / convertGB

	fmt.Printf("%.1fG", ramUsed)
}
