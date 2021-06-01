---
layout:     post
title:      开发 http 的 server 和 client
subtitle:   开发 http 的 server 和 client
date:       2021-05-31
author:     锐玩道
header-img: img/bg_img/post-bg-cook.jpg
catalog:    true
theme:      smartblue
tags:
    - golang
---

> 如果❤️我的文章有帮助，欢迎点赞、关注。这是对我继续技术创作最大的鼓励。[更多往期文章在我的个人博客](https://coderdao.github.io/)

## http 服务器
```go
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/hello", sayHello)

	// 创建服务器
	server := &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	// 监听端口并提供服务
	log.Println("starting httpserver at http:localhost:1210")
	log.Fatal(server.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("hello hello, this is httpserver"))
}
```

启动服务器
```shell
$ go run demo/base/http/server/server.go
2021/05/31 22:26:35 starting httpserver at http:localhost:1210
```

使用 浏览器 或者 命令行测试一下：
```shell
$ curl -v http://localhost:1210/hello

* Trying ::1:1210...
* Connected to localhost (::1) port 1210 (#0)
> GET /hello HTTP/1.1
> Host: localhost:1210
> User-Agent: curl/7.69.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Mon, 31 May 2021 14:28:28 GMT
< Content-Length: 31
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact

hello hello, this is httpserver
```

如上所示：这就是我们服务端返回的内容 `hello hello, this is httpserver`

## http 客户端

### 为什么需要客户端
在多项目、微服务的场景下，项目服务之间的互相通信并不像。使用浏览器、命令行输入域名返回结果。所以需要自己编写发起 http 请求的客户端，实现项目服务之间的通信

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连击池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls 握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue状态码超时时间
	}

	// 创建客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // 没饿
	}

	// 请求数据
	resp, err := client.Get("http://localhost:1210/hello")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取数据
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bds))
}

```

运行服务测试一下`go run demo/base/http/server/server.go`，返回 服务端响应内容 `hello hello, this is httpserver`
