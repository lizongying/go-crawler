---
weight: 3
title: "工具"
---

## 工具

### 证书签名

* -s 自签名服务器证书。如果不设置，会使用本项目默认ca证书进行签名
* -c 新创建ca证书。如果不设置，会使用本项目默认ca证书

开发

```shell
go run tools/tls_generator/*.go
```

构建

```
# 构建
make tls_generator

# 使用
./releases/tls_generator
```

### 中间人代理

```shell
# 默认打印请求和返沪内容
# -f 正则过滤请求
# -p 设置请求代理
# -r 替换返回内容
./releases/mitm

# 测试
# 其他客户端需要信任ca证书 static/tls/ca_crt.pem
curl https://www.baidu.com -x http://localhost:8082 --cacert static/tls/ca.crt
curl https://github.com/lizongying/go-crawler -x http://localhost:8082 --cacert static/tls/ca.crt

```