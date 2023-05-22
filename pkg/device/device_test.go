package device

import (
	"testing"
)

func TestNewDevices(t *testing.T) {
	device, _ := NewDevices("/Users/lizongying/IdeaProjects/go-crawler/pkg/device/devices.csv")
	t.Log(device.Devices)
}
