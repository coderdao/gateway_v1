package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("client failed err:%v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		_, err := conn.Write([]byte("hello abo udp server!"))
		if err != nil {
			fmt.Printf("sell data failed err:%v\n", err)
			return
		}

		retsult := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(retsult)
		if err != nil {
			fmt.Printf("receive data failed, err:%v\n", err)
			return
		}
		fmt.Printf("addr: %v data: %v count: %v\n", remoteAddr, string(retsult[:n]), n)
	}
}
