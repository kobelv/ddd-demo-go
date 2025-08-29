// Package math 提供基础的服务入口 在这里进行初始化操作
package main

import (
	"context"
	"log"

	"ddd-demo-go/interfaces"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := interfaces.NewApp(ctx)
	if err != nil {
		panic(any(err))
	}
	defer app.BeforeShutdown()
	log.Println("server exit:", app.StartServers())
}
