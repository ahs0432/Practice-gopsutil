package main

/*
import (
	"fmt"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

// 가상 메모리 (전체?)
func vMem() {
	vMem, err := mem.VirtualMemory()
	fmt.Print("Virtual Memory: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(vMem)
	}
}

// 스왑 메모리
func sMem() {
	sMem, err := mem.SwapMemory()
	fmt.Print("Swap Memory: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(sMem)
	}
}

// CPU Info
func cpuInfo() {
	info, err := cpu.Info()
	fmt.Print("CPU Info: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(info)
	}
}

// CPU Percent
func cpuPercent() {
	cpuPer, err := cpu.Percent(10, true)
	fmt.Print("CPU Percent: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(cpuPer)
	}
}

// CPU Counts
func cpuCounts() {
	// 논리적인 것도 포함할지 여부를 선택해야할 것으로 확인됨.
	counts, err := cpu.Counts(false)
	fmt.Print("CPU Counts: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(counts)
	}
}

// CPU Times
func cpuTimes() {
	// percpu 뜻을 한번 확인해봐야함.
	times, err := cpu.Times(true)
	fmt.Print("CPU Times: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(times)
	}
}

// Load Average
func loadAvg() {
	load, err := load.Avg()
	fmt.Print("Load Average: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(load)
	}
}

// Disk Partitions
func diskPartitions() {
	// 전체 파티션을 확인할지 여부를 선택함.
	partitions, err := disk.Partitions(true)
	fmt.Print("Disk Partitions: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(partitions)
	}
}

// Disk Usage
func diskUsage() {
	// 대상 경로를 지정하게 됨.
	usage, err := disk.Usage("/")
	fmt.Print("Disk Usage: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(usage)
	}
}

// 전체 프로세스 목록
func processList() {
	processes, err := process.Processes()
	fmt.Print("Process List: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(processes)
	}
}

// 전체 프로세스 PID 목록
func processPID() {
	pids, err := process.Pids()
	fmt.Print("Process PID: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(pids)
	}
}

// 호스트 정보 확인
func hostInfo() {
	info, err := host.Info()
	fmt.Print("Host Info: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(info)
	}
}

// Docker 목록 확인
func dockerIDList() {
	idList, err := docker.GetDockerIDList()
	fmt.Print("Docker ID List: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(idList)
	}
}

// Docker 통계 확인
func dockerStat() {
	stat, err := docker.GetDockerStat()
	fmt.Print("Docker Stat: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(stat)
	}
}

// 네트워크 인터페이스 정보 확인
func netInterfaces() {
	interfaces, err := net.Interfaces()
	fmt.Print("Network Interfaces: ")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(interfaces)
	}
}

func test() {
	vMem()
	sMem()
	cpuInfo()
	cpuPercent()
	cpuCounts()
	cpuTimes()
	loadAvg()
	diskPartitions()
	diskUsage()
	processList()
	processPID()
	hostInfo()
	dockerIDList()
	dockerStat()
	netInterfaces()
}
*/
