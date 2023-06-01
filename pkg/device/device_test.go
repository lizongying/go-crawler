package device

import (
	"github.com/lizongying/go-crawler/static"
	"testing"
)

func TestNewDevices(t *testing.T) {
	devices, _ := NewDevicesFromBytes(static.Devices)
	t.Log(devices.Devices)
}
