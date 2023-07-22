package device

import (
	"bytes"
	"encoding/csv"
	"github.com/lizongying/go-crawler/pkg"
	"log"
	"os"
	"strings"
)

type Device struct {
	Platform    pkg.Platform
	Browser     pkg.Browser
	UserAgent   string
	Fingerprint string
}

type Devices struct {
	Devices map[string][]Device
}

func NewDevicesFromPath(path string) (d *Devices, err error) {
	devices := readCsvFileFromPath(path)
	d = &Devices{
		Devices: devices,
	}

	return
}

func NewDevicesFromBytes(bs []byte) (d *Devices, err error) {
	devices := readCsvFileFromBytes(bs)
	d = &Devices{
		Devices: devices,
	}

	return
}

func readCsvFileFromPath(filePath string) (devices map[string][]Device) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	devices = make(map[string][]Device)
	for _, v := range records {
		device := Device{}
		if len(v) > 1 {
			device.UserAgent = v[1]
		}
		if len(v) > 2 {
			device.Fingerprint = v[2]
		}
		devices[v[0]] = append(devices[v[0]], device)
	}

	return
}

func readCsvFileFromBytes(bs []byte) (devices map[string][]Device) {
	csvReader := csv.NewReader(bytes.NewReader(bs))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}

	devices = make(map[string][]Device)
	for _, v := range records {
		key := strings.Split(v[0], "-")
		device := Device{
			Platform: pkg.Platform(key[0]),
			Browser:  pkg.Browser(key[1]),
		}
		if len(v) > 1 {
			device.UserAgent = v[1]
		}
		if len(v) > 2 {
			device.Fingerprint = v[2]
		}
		devices[v[0]] = append(devices[v[0]], device)
	}

	return
}
