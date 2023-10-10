package networkhardware

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Hardware struct {
	CPUInfo     []cpu.InfoStat         `json:"cpu_info"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memory_info"`
	HostName    string                 `json:"host_name"`
	IPAddresses []string               `json:"ip_addresses"`
}

func HardwareInfo() (*Hardware, error) {
	var info Hardware

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

	// Coletar informações do host (nome da máquina)
	hinfo, err := host.Info()
	if err != nil {
		return nil, err
	}
	info.HostName = hinfo.Hostname

	// Coletar IP da máquina
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, intf := range interfaces {
		for _, addr := range intf.Addrs {
			ips = append(ips, addr.Addr)
		}
	}
	info.IPAddresses = ips

	return &info, nil
}

func HardwareInfoJSON() (string, error) {
	info, err := HardwareInfo()
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
