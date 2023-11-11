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

# spiders
curl "http://127.0.0.1:8090/spiders" -X POST -H "Content-Type: application/json" -H "X-API-Key: 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"

# job run
# once
curl "http://127.0.0.1:8090/job/run" -X POST -d '{"timeout": 2, "name": "test-must-ok", "func": "TestOk", "args": "", "mode": 1}' -H "Content-Type: application/json" -H "X-API-Key: 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
# {"code":0,"msg":"","data":{"id":"133198dc7a0911ee904b9221bc92ca26","start_time":0,"finish_time":0}}

# loop
curl "http://127.0.0.1:8090/job/run" -X POST -d '{"timeout": 2000, "name": "test-must-ok", "func": "TestOk", "args": "", "mode": 2}' -H "Content-Type: application/json" -H "X-API-Key: 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
# {"code":0,"msg":"","data":{"id":"133198dc7a0911ee904b9221bc92ca26","start_time":0,"finish_time":0}}

# job stop
curl "http://127.0.0.1:8090/job/stop" -X POST -d '{"spider_name": "test-must-ok", "job_id": "894a6fe87e2411ee95139221bc92ca26"}' -H "Content-Type: application/json" -H "X-API-Key: 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
# {"code":0,"msg":"","data":{"name":"test-must-ok"}}

# job rerun
curl "http://127.0.0.1:8090/job/rerun" -X POST -d '{"spider_name": "test-must-ok", "job_id": "894a6fe87e2411ee95139221bc92ca26"}' -H "Content-Type: application/json" -H "X-API-Key: 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
# {"code":0,"msg":"","data":{"name":"test-must-ok"}}

```