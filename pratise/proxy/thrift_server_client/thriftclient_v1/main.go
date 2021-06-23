package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/pratise/proxy/thrift_server_client/gen-go/thrift_gen"
	"github.com/coderdao/thrift011/lib/go/thrift"
	"log"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:6001", "input addr")
	flag.Parse()
	if *addr == "" {
		flag.Usage()
		os.Exit(1)
	}
	for {
		tSocket, err := thrift.NewTSocket(*addr)
		if err != nil {
			log.Fatalln("tSocket error:", err)
		}
		transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
		transport, _ := transportFactory.GetTransport(tSocket)
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		client := thrift_gen.NewFormatDataClientFactory(transport, protocolFactory)
		if err := transport.Open(); err != nil {
			log.Fatalln("Error opening:", *addr)
		}
		defer transport.Close()
		data := thrift_gen.Data{Text: "ping"}
		d, err := client.DoFormat(context.Background(), &data)
		if err != nil {
			fmt.Println("err:", err.Error())
		} else {
			fmt.Println("Text:", d.Text)
		}
		time.Sleep(40 * time.Millisecond)
	}
}