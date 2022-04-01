package repository

import (
	model "meshop-product-service/application/model"
	"meshop-product-service/datasource"
)

/**
实现类
*/
type Product struct {
	Base
}

func NewProduct() *Product {
	product := new(Product)
	product.db = datasource.GetDB()
	return product
}

func (this *Product) Find(id uint32) (*model.Product, error) {
	product := &model.Product{}
	product.ID = uint(id)
	if err := this.db.First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (this *Product) Create(product *model.Product) error {
	if err := this.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (this *Product) Update(product *model.Product) (*model.Product, error) {
	if err := this.db.Model(&product).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (this *Product) Delete(product *model.Product) (*model.Product, error) {
	if product.ID > 0 {
		if err := this.db.Delete(model.Product{}, "id=?", product.ID).Error; err != nil {
			return nil, err
		}
	}

	if len(product.Name) > 0 {
		if err := this.db.Delete(model.Product{}, "name=?", product.Name).Error; err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (this *Product) FindListLimit(size int32) ([]model.Product, error) {
	var productList []model.Product
	if err := this.db.Limit(size).Order("id desc").Find(&productList).Error; err != nil {
		return nil, err
	}
	return productList, nil
}

func (this *Product) FindByField(key string, value string, fields string) (*model.Product, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	product := &model.Product{}
	if err := this.db.Select(fields).Where(key+" = ?", value).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
