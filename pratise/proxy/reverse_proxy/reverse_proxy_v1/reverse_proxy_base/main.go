package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

func main()  {
	http.HandleFunc("/", handler)
	log.Println("start serving on port" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	proxy_addr = "http://127.0.0.1:2003"
	port = "2002"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	// 解析代理地址，更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	// 请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil{
		log.Print(err)
		return
	}

	// 把下游请求内容 返回给 上游
	for k, v2header := range resp.Header {
		for _, v := range v2header {
			w.Header().Add(k, v)
		}
	}

	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)
}