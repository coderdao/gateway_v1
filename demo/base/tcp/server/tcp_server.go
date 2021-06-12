package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	for {

		// 建立套字链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		// 创建协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 思考题： 这里不填写会有什么问题

	for {

		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("read from connect fail, err: %v\n", err)
			continue
		}

		str := string(buf[:n])
		fmt.Printf("receive from client, fdata:%v\n", str)
	}
}
