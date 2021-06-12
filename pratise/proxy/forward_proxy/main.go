package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct {
	
} 


func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	
	// 默认连接池
	transport := http.DefaultTransport
	
	// 1. 浅拷贝对象，然后新增 用户真实IP 属性
	outReq := new(http.Request)
	*outReq = *req
	
	// 赋值 “真实ip属性” ;
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err != nil {
		// 如果原始 X-Forwarded-For 有值，追加
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior,", ")+", "+clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}


	// 2.请求下游
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// 3. 把下游请求内容返回给上游
	// 设入 response Header
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)

	// 设入 response 内容
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	fmt.Println("Serve on: 8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
