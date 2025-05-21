package main

import (
	"context"
	"fmt"
	"github.com/gomsr/xxl-job-client"
	"github.com/gomsr/xxl-job-client/logger"
	"github.com/gomsr/xxl-job-client/option"
	"log"
)

func main() {
	client := xxl.NewXxlClient(
		option.WithAppName("xxl-app-test"),
		option.WithClientPort(8080),
		option.WithAdminAddress("http://192.168.191.122:8888/xxl-job-admin/"),
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
