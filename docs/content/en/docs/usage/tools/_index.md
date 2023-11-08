---
weight: 3
title: "Tools"
---

## Tools

### Certificate

```shell
# CA-Signed Certificate
./releases/tls

# Self-Signed Certificate
./releases/tls -s

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