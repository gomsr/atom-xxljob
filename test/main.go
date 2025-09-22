package main

import (
	"context"
	"fmt"
	"github.com/gomsr/atom-xxljob"
	"github.com/gomsr/atom-xxljob/configx"
	"github.com/gomsr/atom-xxljob/logger"
	"log"
)

func main() {
	client := xxl.NewXxlClientOps(
		configx.WithAppName("xxl-app-test"),
		configx.WithClientPort(8080),
		configx.WithAdminAddress("http://192.168.191.122:8888/xxl-job-admin/"),
	)
	defer func() {
		client.ExitApplication()
		_ = client.Close()
	}()
	client.RegisterJob("HelloWorld", HelloWorld)
	if err := client.Run(); err != nil {
		log.Println(err)
	}
}

func HelloWorld(ctx context.Context) error {
	for i := 0; i < 100; i++ {
		logger.Info(ctx, fmt.Sprintf("hello world:%d", i))
	}
	return nil
}
