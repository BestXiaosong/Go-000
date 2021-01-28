package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		reflect.TypeOf(conn)
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		conn.Write([]byte("xiaosong"))
		fmt.Printf("Received data: %v\n", string(buf[:n]))
	}
}
