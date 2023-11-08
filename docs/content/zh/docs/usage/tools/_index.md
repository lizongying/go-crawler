---
weight: 3
title: "工具"
---

## 工具

### 证书签名

```shell
# 证书签名
./releases/tls

# 自签名
./releases/tls -s
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