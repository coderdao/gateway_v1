package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 读取命令行输入，建立套接字链接

// 创建处理协程

func main() {
	// 链接服务器，监听端口
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()

	if err != nil {
		fmt.Printf("connect failed, err:%v\n", err.Error())
	}
	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 一直读取直到读取到 \n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err.Error())
			break
		}

		// 读取到Q时停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}

		// 回复服务器消息
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err.Error())
			break
		}
	}

}
