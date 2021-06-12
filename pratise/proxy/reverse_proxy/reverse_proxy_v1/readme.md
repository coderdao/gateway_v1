## 方向代理demo

### 启动 真实服务器
```
$go run /d/Dev/workplace/golang/gateway_v1/pratise/proxy/reverse_proxy/real_server/main.go
2021/06/12 17:59:24 starting httpserver at 127.0.0.1:2003
2021/06/12 17:59:24 starting httpserver at 127.0.0.1:2004
```

## 测试 真实服务器
```bash
$ curl 'http://127.0.0.1:2003/sda?sda=111'
http://127.0.0.1:2003/sda
RemoteAddr=127.0.0.1:49795,X-Forwarded-For=,X-Real-Ip=
headers =map[Accept:[*/*] User-Agent:[curl/7.69.1]]


$ curl 'http://127.0.0.1:2004/sda?sda=111'
http://127.0.0.1:2003/sda
RemoteAddr=127.0.0.1:49795,X-Forwarded-For=,X-Real-Ip=
headers =map[Accept:[*/*] User-Agent:[curl/7.69.1]]
```

### 启动 代理服务器
```
$go run /d/Dev/workplace/golang/gateway_v1/pratise/proxy/reverse_proxy/reverse_proxy_base/main.go
2021/06/12 18:32:21 start serving on port2002
```

## 测试 代理服务器
```bash
$ curl 'http://127.0.0.1:2002/sda?sda=111'

# 转发至 2003 并返回
http://127.0.0.1:2003/sda
RemoteAddr=127.0.0.1:57595,X-Forwarded-For=,X-Real-Ip=
headers =map[Accept:[*/*] Accept-Encoding:[gzip] User-Agent:[curl/7.69.1]]

```