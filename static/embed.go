package static

import (
	"embed"
	_ "embed"
)

//go:embed devices.csv
var Devices []byte

//go:embed tls/ca_crt.pem
var CaCert []byte

//go:embed tls/ca_key.pem
var CaKey []byte

//go:embed tls/server_crt.pem
var ServerCert []byte

//go:embed tls/server_key.pem
var ServerKey []byte

//go:embed tls/server_self_crt.pem
var ServerSelfCert []byte

//go:embed tls/server_self_key.pem
var ServerSelfKey []byte

//go:embed statics
var Statics embed.FS

//go:embed html
var Html embed.FS

//go:embed api
var Api embed.FS
