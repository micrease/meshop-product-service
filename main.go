package main

import (
	"context"
	"flag"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/micrease/micrease-core/errs"
	nacos "github.com/micrease/micrease-core/registry"
	"meshop-product-service/application/handler"
	sysConfig "meshop-product-service/config"
	"meshop-product-service/datasource"
)

func InitServer() {
	//解析命令运行参数
	flag.Parse()
	//从配置中心获取业务配置
	sysConfig.InitSysConfig()
	//连接数据库
	datasource.InitDatabase()
}

func RecoverWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		defer errs.Recover(&err)
		err = fn(ctx, req, rsp)
		return err
	}
}

func main() {
	InitServer()
	conf := sysConfig.Get()
	log.Info("Version:", conf.Service.Version)
	//注册
	registry := nacos.NewRegistry(registry.Addrs(conf.Nacos.Addrs))

	opts := []micro.Option{
		micro.Address(":" + conf.Service.Port),
		micro.Name(conf.Service.ServiceName),
		micro.Version(conf.Service.Version),
		micro.Registry(registry),
		micro.WrapHandler(RecoverWrapper),
	}
	svr := micro.NewService(opts...)

	svr.Init()

	//注册grpc handler
	handler.RegisterProduct(svr)
	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
