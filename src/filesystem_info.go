package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func convertToGB(number uint64) (uint64) {
	return number / 1e9
}

// return used and full memory
func getFsInfo(path string) (uint64, uint64) {
	var statFs unix.Statfs_t

	var err error = unix.Statfs(path, &statFs)
	if err != nil {
		return 0, 0
	}

	// the conversion is necessary because Bsize isn't retrieved as uin64 but int64
	var totalStorage uint64 = statFs.Blocks * uint64(statFs.Bsize)
	var freeStorage uint64 = statFs.Bfree *  uint64(statFs.Bsize) // returns the 'free' memory, it includes also restricted memory for the OS

	var usedStorage uint64 = totalStorage - freeStorage

	return usedStorage, totalStorage
}

func main() {
	var filePath string = "/"

	var usedStr, totalStr uint64 = getFsInfo(filePath)
	var percentageUsed uint8 = uint8((float64(usedStr) / float64(totalStr)) * 100)

	fmt.Printf("%03d/%03d (%02d%%)", convertToGB(usedStr), convertToGB(totalStr), percentageUsed)
}
