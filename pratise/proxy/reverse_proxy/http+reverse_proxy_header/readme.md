# 代理转发 - 修改请求头

说明文件： _md/3.11 实现 ReverseProxy Header 头.md


## 转发路径
reverse_proxy_level1（监听 2001 - 转发 2002）
reverse_proxy_level2（监听 2002 - 转发 2003）
real_server（监听 2003）


## 测试
首先单独请求 2001
> curl -H 'X-Forwarded-For: 2.2.2.2' '127.0.0.1:2001/test'
```bash
$ curl -H 'X-Forwarded-For: 2.2.2.2' '127.0.0.1:2001/test'

http://127.0.0.1:2004/test
RemoteAddr=127.0.0.1:56197,X-Forwarded-For=2.2.2.2, 127.0.0.1, 127.0.0.1,X-Real-Ip=127.0.0.1:56195
headers =map[Accept:[*/*] Accept-Encoding:[gzip] User-Agent:[curl/7.69.1] X-Forwarded-For:[2.2.2.2, 127.0.0.1, 127.0.0.1] X-Real-Ip:[127.0.0.1:56195]]


```