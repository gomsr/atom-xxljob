# xxl-job-client

## xxl-job go sdk
兼容 xxl-job v2.4.0

## demo

```go
package main

import (
	"context"
	"fmt"
	xxl "github.com/gomsr/atom-xxljob"
	"github.com/gomsr/atom-xxljob/logger"
	"github.com/gomsr/atom-xxljob/option"
	"log"
)

func main() {
	client := xxl.NewXxlClient(
		option.WithAppName("执行器的名字"),
		option.WithClientPort(8080),
		option.WithAdminAddress("xxl-job接入地址"),
	)
	defer func() {
		client.ExitApplication()
		client.Close()
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
```