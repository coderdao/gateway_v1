package main

import (
	middleware "gateway/pratise/proxy/middlewarw"
	"gateway/pratise/proxy/proxy"
	"log"
	"net/http"
	"net/url"
)

var addr = "127.0.0.1:2002"

// 熔断方案
func main() {
	// 设置限流配置
	coreFunc := func(c *middleware.SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	sliceRouter := middleware.NewSliceRouter()
	sliceRouter.Group("/").Use(middleware.RateLimiter())
	// 注入到中间件
	routerHandler := middleware.NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}