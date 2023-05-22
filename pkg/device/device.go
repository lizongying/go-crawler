package device

import (
	"encoding/csv"
	"log"
	"os"
)

type Device struct {
	UserAgent string
}

type Devices struct {
	Devices map[string][]Device
}

func NewDevices(filePath string) (d *Devices, err error) {
	devices := readCsvFile(filePath)
	d = &Devices{
		Devices: devices,
	}

	return
}

func readCsvFile(filePath string) (devices map[string][]Device) {
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
		devices[v[0]] = append(devices[v[0]], Device{
			UserAgent: v[1],
		})
	}

	return
}
