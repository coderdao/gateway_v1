package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	// 设置真实服务器
	rs1 := &RealServer{Addr:"127.0.0.1:2003"}
	rs1.Run()

	rs2 := &RealServer{Addr:"127.0.0.1:2004"}
	rs2.Run()

	// 监听服务关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

/** 服务器设置 */
type RealServer struct {
	Addr string
}

func (r *RealServer) Run()  {
	log.Println("starting httpserver at "+r.Addr)

	// 启动路由，设置路由规则 => 对应处理方法
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error", r.ErrorHandler)
	mux.HandleFunc("/base/timeout", r.TimeoutHandler)

	server := &http.Server{
		Addr:              r.Addr,
		WriteTimeout:      time.Second * 3,
		Handler: 			mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	/**
	127.0.0.1:8080/abc?sdsdsa=11
	r.Addr=127.0.0.1:8080
	req.URL.Path=/abc
	 */

	// 转发前, 设置请求机器参数
	upath := fmt.Sprintf("http://%s%s\n", r.Addr, req.URL.Path) // 打印请求地址
	realIP := fmt.Sprintf(
		"RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v\n",
		req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))
	header:=fmt.Sprintf("headers =%v\n",req.Header)

	// 写入返回体 response
	io.WriteString(w, upath)
	io.WriteString(w, realIP)
	io.WriteString(w, header)
}

func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	w.WriteHeader(500)

	io.WriteString(w, upath)
}

func (r *RealServer) TimeoutHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(6*time.Second)
	upath := "timeout handler"
	w.WriteHeader(200)
	io.WriteString(w, upath)
}