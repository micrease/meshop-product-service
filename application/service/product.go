package service

import (
	"github.com/micrease/meshop-protos/product/pb"
	"github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/errs"
	"meshop-product-service/application/repo"
)

type Product struct {
	repo *repo.Product
}

func NewProduct() *Product {
	product := &Product{}
	product.repo = repo.NewProduct()
	return product
}

func (this *Product) List(ctx *context.Context, req *pb.ProductRequest, resp *pb.ProductListResponse) error {
	products, err := this.repo.Get(req.Size)

	errs.PanicIf(err, 5001, "查询商品列表失败")

	//resp.Data = products
	prods := make([]*pb.Product, 0)
	var i int
	for i = 0; i < len(products); i++ {
		prods = append(prods, &pb.Product{ProdId: int32(products[i].ID), ProdName: products[i].Name})
	}
	resp.Data = prods
	return nil
}

/**
 * 获取详情
 */
func (this *Product) Detail(ctx *context.Context, req *pb.Product, resp *pb.ProductResponse) error {
	//prod :=  &model.Product{ID: uint(req.ProdId), Name: req.ProdName}
	if prod, err := this.repo.GetOne(uint32(req.ProdId)); err != nil {
		return errs.Err(5002, "获取商品详情失败", err)
	} else {
		resp.Data = &pb.Product{ProdId: int32(prod.ID), ProdName: prod.Name}
	}
	return nil
}
