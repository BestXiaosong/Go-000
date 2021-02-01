package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func main() {

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return StartTcp(ctx, ":8080")
	})

	group.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-signals:
			fmt.Println("receive quit signal")
			return errors.New("receive quit signal")
		case <-ctx.Done():
			fmt.Println("signal ctx done")
			return ctx.Err()
		}
	})

	err := group.Wait() // first error return
	fmt.Println("group err:", err)
}

func StartTcp(ctx context.Context, addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil // 终止程序
	}

	go func() {
		select {
		case <-ctx.Done():
			_ = listener.Close()
		}
	}()
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		reflect.TypeOf(conn)
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return err // 终止程序
		}
		go doServerStuff(ctx, conn)
	}
}

func doServerStuff(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	msgChan := make(chan []byte, 1)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msgChan <- input.Bytes()
	}

	go sendMessage(ctx, conn, msgChan)

}

func sendMessage(ctx context.Context, conn net.Conn, ch chan []byte) {
	select {
	case <-ctx.Done():
		return
	case msg := <-ch:
		fmt.Println(conn, string(msg))
	}
}
