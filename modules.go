package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Virtual Memory 임계값에 따른 구현
func vMemCheck(threshold float64, types bool) bool {
	vMem, err := mem.VirtualMemory()

	if err != nil {
		panic(err)
	}

	// 만약 임계값이 %인지 아니면 총 사용량인지에 따름
	if types {
		if vMem.UsedPercent >= threshold {
			return false
		}
	} else {
		// 값은 MB로 비교하도록 함
		if vMem.Used >= (uint64(threshold) * 1024 * 1024) {
			return false
		}
	}

	return true
}

// Swap Memory 임계값에 따른 구현
func sMemCheck(threshold float64, types bool) bool {
	vMem, err := mem.SwapMemory()

	if err != nil {
		panic(err)
	}

	// 만약 임계값이 %인지 아니면 총 사용량인지에 따름
	if types {
		if vMem.UsedPercent >= threshold {
			return false
		}
	} else {
		// 값은 MB로 비교하도록 함
		if vMem.Used >= (uint64(threshold) * 1024 * 1024) {
			return false
		}
	}

	return true
}

// 대상 경로의 사용량을 임계값에 따른 구현
func diskUsedCheck(threshold float64, types bool, path string) bool {
	usage, err := disk.Usage(path)

	if err != nil {
		panic(nil)
	}

	// 만약 임계값이 %인지 아니면 총 사용량인지에 따름
	if types {
		if usage.UsedPercent >= threshold {
			return false
		}
	} else {
		// 값은 MB로 비교하도록 함
		if usage.Used >= (uint64(threshold) * 1024 * 1024) {
			return false
		}
	}

	return true
}

// 대상 경로의 inode 임계값에 따른 구현
func diskinodeCheck(threshold float64, types bool, path string) bool {
	usage, err := disk.Usage(path)

	if err != nil {
		panic(nil)
	}

	// 만약 임계값이 %인지 아니면 총 사용량인지에 따름
	if types {
		if usage.InodesUsedPercent >= threshold {
			return false
		}
	} else {
		// 값은 MB로 비교하도록 함
		if usage.InodesUsed >= uint64(threshold) {
			return false
		}
	}

	return true
}

// CPU 사용량 임계값에 따른 구현
func cpuPerCheck(threshold float64) bool {
	cpuPer, err := cpu.Percent(0, false)

	if err != nil {
		panic(nil)
	}

	// 평균 값에 따른 비교
	if cpuPer[0] >= threshold {
		return false
	}

	return true
}
