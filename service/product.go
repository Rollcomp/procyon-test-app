package service

import (
	configure "github.com/procyon-projects/procyon-configure"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
	web "github.com/procyon-projects/procyon-web"
)

type ProductService interface {
	FindAll(ctx *web.WebRequestContext) []*model.Product
	FindById(ctx *web.WebRequestContext, id int) *model.Product
	Save(ctx *web.WebRequestContext, product *model.Product) *model.Product
	Update(ctx *web.WebRequestContext, id int, updatedProduct *model.Product) *model.Product
	DeleteById(ctx *web.WebRequestContext, id int)
}

type ImpProductService struct {
	productRepository    repository.ProductRepository
	transactionalContext tx.TransactionalContext
}

func NewProductService(productRepository repository.ProductRepository,
	transactionalContext tx.TransactionalContext, properties configure.WebServerProperties) ImpProductService {
	return ImpProductService{
		productRepository,
		transactionalContext,
	}
}

func (service ImpProductService) FindAll(ctx *web.WebRequestContext) []*model.Product {
	/*products, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.FindAll(ctx), nil
	})
	return products.([]*model.Product)*/
	return []*model.Product{
		{
			Name:     "Product-1",
			Category: 8,
		},
		{
			Name:     "Product-2",
			Category: 7,
		},
	}
}

func (service ImpProductService) FindById(ctx *web.WebRequestContext, id int) *model.Product {
	/*result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		return product, nil
	})
	return result.(*model.Product)*/
	return &model.Product{
		Name:     "Test Product",
		Category: 1,
	}
}

func (service ImpProductService) Save(ctx *web.WebRequestContext, product *model.Product) *model.Product {
	/*result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.Save(ctx, product), nil
	})*/
	return product
}

func (service ImpProductService) Update(ctx *web.WebRequestContext, id int, updatedProduct *model.Product) *model.Product {
	/*result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		return service.productRepository.Update(ctx, updatedProduct), nil
	})*/
	return updatedProduct
}

func (service ImpProductService) DeleteById(ctx *web.WebRequestContext, id int) {
	/*service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		service.productRepository.DeleteById(ctx, id)
		return nil, nil
	})*/
}
