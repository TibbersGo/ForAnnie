package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

func main() {

	//创建app
	app := hydra.NewApp(
		hydra.WithPlatName("test"),        //平台名
		hydra.WithSystemName("apiserver"), //系统或应用名
		hydra.WithServerTypes(http.API),
	)

	//注册服务
	app.API("/hello", func(ctx hydra.IContext) interface{} {
		return "hello world"
	})

	//启动app
	app.Start()
}
