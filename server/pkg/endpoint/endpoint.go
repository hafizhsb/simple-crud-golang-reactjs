package endpoint

import (
	"context"
	"stock/pkg/entity"
	"stock/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	Status        endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	GetProduct    endpoint.Endpoint
	CreateProduct endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
}

func MakeServerEndpoints(s service.Service) Endpoint {
	return Endpoint{
		Status:        makeStatus(s),
		GetProducts:   makeGetProducts(s),
		GetProduct:    makeGetProduct(s),
		CreateProduct: makeCreateProduct(s),
		UpdateProduct: makeUpdateProduct(s),
		DeleteProduct: makeDeleteProduct(s),
	}
}

func makeStatus(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.Status(ctx, request)
	}
}

func makeGetProducts(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetProduct)
		return s.GetProducts(ctx, req)
	}
}

func makeGetProduct(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DetailProduct)
		return s.GetProduct(ctx, req)
	}
}

func makeCreateProduct(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DetailProduct)
		return s.CreateProduct(ctx, req)
	}
}

func makeUpdateProduct(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DetailProduct)
		return s.UpdateProduct(ctx, req)
	}
}

func makeDeleteProduct(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetProduct)
		return s.DeleteProduct(ctx, req)
	}
}
