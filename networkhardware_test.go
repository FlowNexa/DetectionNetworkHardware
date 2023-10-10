package networkhardware

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHardwareInfo(t *testing.T) {
	info, err := HardwareInfo()

	assert.NoError(t, err, "Failed to get hardware info")
	assert.NotEmpty(t, info.CPUInfo, "Expected non-empty CPUInfo")
	assert.NotNil(t, info.MemoryInfo, "Expected non-nil MemoryInfo")
	assert.True(t, info.MemoryInfo.Total > 0, "Expected positive total memory")
	assert.NotEmpty(t, info.HostName, "Expected non-empty HostName")
	assert.NotEmpty(t, info.IPAddresses, "Expected non-empty IPAddresses")
}

func TestGetHardwareInfoJSON(t *testing.T) {
	jsonStr, err := HardwareInfoJSON()
	assert.NoError(t, err, "Failed to get hardware info as JSON")
	assert.NotEmpty(t, jsonStr, "Expected non-empty JSON string")

	var result Hardware
	err = json.Unmarshal([]byte(jsonStr), &result)
	assert.NoError(t, err, "Failed to unmarshal JSON")
	assert.NotEmpty(t, result.CPUInfo, "Expected non-empty CPUInfo in JSON")
}
