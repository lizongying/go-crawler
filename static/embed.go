package static

import (
	_ "embed"
)

//go:embed devices.csv
var Devices []byte

//go:embed tls/cert.pem
var Cert []byte

//go:embed tls/key.pem
var Key []byte
