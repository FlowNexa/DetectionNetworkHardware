package networkhardware

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type HardwareInfo struct {
	CPUInfo        []cpu.InfoStat             `json:"cpu_info"`
	MemoryInfo     *mem.VirtualMemoryStat     `json:"memory_info"`
	DiskPartitions []disk.PartitionStat       `json:"disk_partitions"`
	DiskUsage      map[string]*disk.UsageStat `json:"disk_usage"`
	NetworkInfo    []net.InterfaceStat        `json:"network_info"`
	HostInfo       *host.InfoStat             `json:"host_info"`
}

func GetHardwareInfo() (*HardwareInfo, error) {
	var info HardwareInfo

	// Coletar informações da CPU
	cpus, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	info.CPUInfo = cpus

	// Coletar informações da memória
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	info.MemoryInfo = vmem

	// Coletar informações dos discos
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	info.DiskPartitions = partitions

	diskUsageMap := make(map[string]*disk.UsageStat)
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			return nil, err
		}
		diskUsageMap[p.Mountpoint] = usage
	}
	info.DiskUsage = diskUsageMap

	// Coletar informações da rede
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	info.NetworkInfo = interfaces

	// Coletar informações do host
	hinfo, err := host.Info()
	if err != nil {
		return nil, err
	}
	info.HostInfo = hinfo

	return &info, nil
}

func GetHardwareInfoJSON() (string, error) {
	info, err := GetHardwareInfo()
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
