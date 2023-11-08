---
weight: 1
title: Api
---

## Api

```shell
go run cmd/multi_spider/*.go -c example.yml
```

```shell
# index
curl "http://127.0.0.1:8080" -H "Content-Type: application/json"

# spider run
curl "http://127.0.0.1:8080/job/run" -X POST -d '{"timeout": 2, "name": "test-must-ok", "func":"TestOk", "args":"", "mode": 1}' -H "Content-Type: application/json"
# {"code":0,"msg":"","data":{"name":"test-must-ok"}}

# spider stop
curl "http://127.0.0.1:8080/job/stop" -X POST -d '{"id":""}' -H "Content-Type: application/json"
# {"code":0,"msg":"","data":{"name":"test-must-ok"}}

```