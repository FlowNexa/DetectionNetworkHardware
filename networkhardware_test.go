package networkhardware

import (
	"testing"
)

func TestGetHardwareInfo(t *testing.T) {
	info, err := GetHardwareInfo()
	if err != nil {
		t.Fatalf("Failed to get hardware info: %v", err)
	}

	if len(info.CPUInfo) == 0 {
		t.Errorf("Expected non-empty CPUInfo")
	}

	if info.MemoryInfo == nil {
		t.Errorf("Expected non-nil MemoryInfo")
	}

	if len(info.DiskPartitions) == 0 {
		t.Errorf("Expected non-empty DiskPartitions")
	}

	if len(info.DiskUsage) == 0 {
		t.Errorf("Expected non-empty DiskUsage")
	}

	if len(info.NetworkInfo) == 0 {
		t.Errorf("Expected non-empty NetworkInfo")
	}

	if info.HostInfo == nil {
		t.Errorf("Expected non-nil HostInfo")
	}
}

func TestGetHardwareInfoJSON(t *testing.T) {
	jsonStr, err := GetHardwareInfoJSON()
	if err != nil {
		t.Fatalf("Failed to get hardware info as JSON: %v", err)
	}

	if len(jsonStr) == 0 {
		t.Errorf("Expected non-empty JSON string")
	}
}
