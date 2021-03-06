package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	stopSignal := make(chan struct{})

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		server := http.Server{
			Addr:    ":8080",
			Handler: nil,
		}
		go func() {
			// 接收到 errGroup.Done 时终止http服务 并发送信号 已结束httpServer
			<-ctx.Done()
			fmt.Printf("http server %s ctx done\n", server.Addr)
			if err := server.Shutdown(context.Background()); err != nil {
				fmt.Printf("http server %s shutdown err :%s", server.Addr, err)
			}

			stopSignal <- struct{}{}

		}()
		return server.ListenAndServe()
	})

	group.Go(func() error {
		server := http.Server{
			Addr:    ":8081",
			Handler: nil,
		}
		go func() {
			// 接收到 errGroup.Done 时终止http服务 并发送信号 已结束httpService
			<-ctx.Done()
			fmt.Printf("http server %s ctx done \n", server.Addr)
			if err := server.Shutdown(context.Background()); err != nil {
				fmt.Printf("http server %s shutdown err :%s", server.Addr, err)
			}
			stopSignal <- struct{}{}
		}()
		return server.ListenAndServe()
	})

	// 监听系统信号
	group.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		// 接收到终止信号 返回错误终止运行
		select {
		case <-signals:
			fmt.Println("receive quit signal")
			return errors.New("receive quit signal")
		case <-ctx.Done():
			fmt.Println("signal ctx done")
			return ctx.Err()
		}

	})

	fmt.Println("main running")

	if err := group.Wait(); err != nil {
		fmt.Println("err group wait err:", err.Error())
	}

	// 启动了两个httpServer 都结束后终止
	for i := 0; i < 2; i++ {
		<-stopSignal
	}

	fmt.Println("all stopped!")
}
