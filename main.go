package main

import (
	"context"
	"flag"
	"github.com/micrease/micrease-core/config"
	nacos "github.com/micrease/micrease-core/registry"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"github.com/micro/go-plugins/wrapper/service/v2"
	"meshop-product-service/application/handler"
	sysConfig "meshop-product-service/config"
	"meshop-product-service/datasource"
)

func InitServer() {
	//解析命令运行参数
	flag.Parse()
	//从config.json加载配置信息
	log.Info("Config ResourcesPath:", *config.ResourcesPath)
	log.Info("Config Env:", *config.Env)
	//从配置中心获取业务配置
	sysConfig.InitSysConfig()
	//连接数据库
	datasource.InitDatabase()
}

func NewLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		return err
	}
}

func main() {
	InitServer()
	conf := sysConfig.Get()
	log.Info("Version:", conf.Service.Version)
	//注册中心
	registry := nacos.NewRegistry(registry.Addrs(conf.Nacos.Addrs))
	//负载均衡策略:轮循
	wrapper := roundrobin.NewClientWrapper()
	opts := []micro.Option{
		micro.Address(":" + conf.Service.Port),
		micro.Name(conf.Service.ServiceName),
		micro.Version(conf.Service.Version),
		micro.Registry(registry),
		micro.WrapClient(wrapper),
		micro.WrapHandler(NewLogWrapper),
	}
	svr := micro.NewService(opts...)

	svr.Init(
		//把服务自身注册进去，在handler中可以获取自身信息
		micro.WrapClient(service.NewClientWrapper(svr)),
		micro.WrapHandler(service.NewHandlerWrapper(svr)),
	)

	//注册grpc handler
	handler.RegisterProduct(svr)
	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
