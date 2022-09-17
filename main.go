package main

import "fmt"

func main() {
	//test()

	// 아래 구현된 것은 false가 임계치 이상 작동된 것
	// Virtual Memory 90% 이상 사용 여부
	fmt.Println(vMemCheck(90.0, true))

	// Virtual Memory 1024 MB(1 GB) 이상 사용 여부
	fmt.Println(vMemCheck(1024.0, false))

	// Swap Memory 90% 이상 사용 여부
	fmt.Println(sMemCheck(90.0, true))

	// Swap Memory 1024 MB(1 GB) 이상 사용 여부
	fmt.Println(sMemCheck(1024.0, false))

	// Disk 사용량 90% 이상 사용 여부
	fmt.Println(diskUsedCheck(90.0, true, "/"))

	// Disk 사용량 102400 MB(100 GB) 이상 사용 여부
	fmt.Println(diskUsedCheck(102400.0, false, "/"))

	// Disk inode 사용량 50% 이상 사용 여부
	fmt.Println(diskinodeCheck(50.0, true, "/"))

	// Disk inode 사용량 1000개 이상 사용 여부
	fmt.Println(diskinodeCheck(1000.0, false, "/"))

	// CPU 사용량 80% 이상 사용 여부
	fmt.Println(cpuPerCheck(80.0))
}
