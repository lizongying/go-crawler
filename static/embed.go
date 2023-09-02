package static

import (
	"embed"
	_ "embed"
)

//go:embed devices.csv
var Devices []byte

//go:embed tls/cert.pem
var Cert []byte

//go:embed tls/key.pem
var Key []byte

//go:embed statics
var Statics embed.FS

//go:embed html
var Html embed.FS
