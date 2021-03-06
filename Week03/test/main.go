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

	var stopSignal chan struct{}

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		server := http.Server{
			Addr:    ":8080",
			Handler: nil,
		}
		go func() {
			<-ctx.Done()
			fmt.Println("http server 8080 ctx done")
			if err := server.Shutdown(context.Background()); err != nil {
				fmt.Println("http server 8080 shutdown err :", err)
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
			<-ctx.Done()
			fmt.Println("http server 8081 ctx done")
			if err := server.Shutdown(context.Background()); err != nil {
				fmt.Println("http server 8081 shutdown err :", err)
			}
			stopSignal <- struct{}{}
		}()
		return server.ListenAndServe()
	})

	// 监听系统信号
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

	fmt.Println("main running")

	if err := group.Wait(); err != nil {
		fmt.Println("err group wait err:", err.Error())
	}
	<-stopSignal

	fmt.Println("all stopped!")
}
