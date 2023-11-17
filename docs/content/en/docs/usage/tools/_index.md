---
weight: 3
title: "Tools"
---

## Tools

### Certificate

* `-s` Self-signed server certificate. If not set, the default CA certificate of this project will be used for signing.
* `-c` Create a new CA certificate. If not set, the default CA certificate of this project will be used.

dev

```shell
go run tools/tls_generator/*.go
```

build

```
# build
make tls_generator

# run
./releases/tls_generator
```

### MITM

```shell
# Print request and response by default
# -f Filter requests using regular expressions.
# -p Set request proxy.
# -r Replace the response
./releases/mitm

# Test
# Other clients need to trust the CA certificate. static/tls/ca_crt.pem
curl https://www.baidu.com -x http://localhost:8082 --cacert static/tls/ca.crt
curl https://github.com/lizongying/go-crawler -x http://localhost:8082 --cacert static/tls/ca.crt

```