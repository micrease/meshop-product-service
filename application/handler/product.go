package handler

import (
	"context"
	"github.com/micrease/meshop-protos/product/pb"
	"github.com/micrease/micrease-core/errs"
	"github.com/micro/go-micro/v2"
	micro_errors "github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
	"meshop-product-service/application/model"
	"meshop-product-service/application/repo"
	"meshop-product-service/application/service"
)

type Product struct {
	//初始化上下文传递的通用变量，比如log,gorm,redis等
	RpcHandler
	//初始化
	service *service.Product
}

//把自身注册到rpc handler中
func RegisterProduct(svr micro.Service) {
	s := &Product{}
	//1,创建上下文
	s.NewContext()
	s.service = service.NewProduct()
	//2,把当前对象注册到rpc中
	pb.RegisterProductServiceHandler(svr.Server(), s)
}

/**
 * 获取列表
 */
func (this *Product) GetProductList(ctx context.Context, req *pb.ProductRequest, resp *pb.ProductListResponse) (err error) {
	defer errs.Recover(&err)
	return this.service.List(this.ctx, req, resp)
}

/**
 * 获取详情
 */
func (this *Product) GetProductDetail(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	return this.service.Detail(this.ctx, req, resp)
}

/**
 * 创建商品
 */
func (this *Product) CreateProduct(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	newProd := model.Product{Name: req.ProdName}
	if err := repo.NewProduct().Create(&newProd); err != nil {
		return errs.Err(5003, "创建商品失败", err)
	}
	return nil
}

/**
 * 更新商品
 */
func (this *Product) UpdateProduct(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	log.Info("CreateProduct..", req)
	newProd := &model.Product{ID: uint(req.ProdId), Name: req.ProdName}
	if prod, err := repo.NewProduct().Update(newProd); err != nil {
		return errs.Err(5004, "更新商品失败", err)
	} else {
		resp.Data = &pb.Product{ProdId: int32(prod.ID), ProdName: prod.Name}
	}
	return nil
}

/**
 * 删除商品
 */
func (this *Product) DeleteProduct(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	log.Info("DeleteProduct..", req)
	delProd := &model.Product{ID: uint(req.ProdId), Name: req.ProdName}
	if prod, err := repo.NewProduct().Delete(delProd); err != nil {
		return errs.Err(5005, "删除商品失败", err)
	} else {
		resp.Data = &pb.Product{ProdId: int32(prod.ID), ProdName: prod.Name}
	}
	return nil
}

func (this *Product) Test(ctx context.Context, req *pb.ProductRequest, resp *pb.ProductListResponse) error {
	service, ok := micro.FromContext(ctx)
	if !ok {
		return micro_errors.InternalServerError("com.example.srv.foo", "Could not retrieve service")
	}
	//panic的时候也可以抛出异常，调用方将会从error中获取到panic信息，但是go micro log中会记录错误信息
	panic("xxxx" + service.Name())
}
