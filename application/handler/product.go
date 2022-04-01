package handler

import (
	"context"
	"github.com/micrease/meshop-protos/product/pb"
	"github.com/micrease/micrease-core/rpc"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
	"meshop-product-service/application/model"
	"meshop-product-service/application/repository"
)

type Product struct {
	rpc.ServiceHandler
}

//创建对象，并把自身注册到handler中
func RegisterProduct(service micro.Service) *Product {
	hdl := new(Product)
	pb.RegisterProductServiceHandler(service.Server(), hdl)
	return hdl
}

func (this *Product) Test(ctx context.Context, req *pb.ProductRequest, resp *pb.ProductListResponse) error {
	service, ok := micro.FromContext(ctx)
	if !ok {
		return errors.InternalServerError("com.example.srv.foo", "Could not retrieve service")
	}
	//panic的时候也可以抛出异常，调用方将会从error中获取到panic信息，但是go micro log中会记录错误信息
	panic("xxxx" + service.Name())
}

/**
 * 获取列表
 */
func (this *Product) GetProductList(ctx context.Context, req *pb.ProductRequest, resp *pb.ProductListResponse) error {
	if products, err := repository.NewProduct().FindListLimit(req.Size); err != nil {
		return this.Error(5001, "查询商品列表失败", err)
	} else {
		prods := make([]*pb.Product, 0)
		var i int
		for i = 0; i < len(products); i++ {
			prods = append(prods, &pb.Product{ProdId: int32(products[i].ID), ProdName: products[i].Name})
		}
		resp.Data = prods
	}

	//此处固定返回nil，否则调用方无法收到返回
	return nil
}

/**
 * 获取详情
 */
func (this *Product) GetProductDetail(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	log.Info("GetProdDetail..", req)
	//prod :=  &model.Product{ID: uint(req.ProdId), Name: req.ProdName}
	if prod, err := repository.NewProduct().Find(uint32(req.ProdId)); err != nil {
		return this.Error(5002, "获取商品详情失败", err)
	} else {
		resp.Data = &pb.Product{ProdId: int32(prod.ID), ProdName: prod.Name}
	}
	return nil
}

/**
 * 创建商品
 */
func (this *Product) CreateProduct(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	newProd := model.Product{Name: req.ProdName}
	if err := repository.NewProduct().Create(&newProd); err != nil {
		return this.Error(5003, "创建商品失败", err)
	}
	return nil
}

/**
 * 更新商品
 */
func (this *Product) UpdateProduct(ctx context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	log.Info("CreateProduct..", req)
	newProd := &model.Product{ID: uint(req.ProdId), Name: req.ProdName}
	if prod, err := repository.NewProduct().Update(newProd); err != nil {
		return this.Error(5004, "更新商品失败", err)
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
	if prod, err := repository.NewProduct().Delete(delProd); err != nil {
		return this.Error(5005, "删除商品失败", err)
	} else {
		resp.Data = &pb.Product{ProdId: int32(prod.ID), ProdName: prod.Name}
	}
	return nil
}
