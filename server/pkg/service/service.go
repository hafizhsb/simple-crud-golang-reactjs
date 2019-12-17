package service

import (
	"context"
	"stock/pkg/entity"
)

type Service interface {
	Status(ctx context.Context, req interface{}) (entity.ResponseHTTP, error)
	GetProducts(ctx context.Context, req entity.GetProduct) (entity.ResponseHTTP, error)
	GetProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error)
	CreateProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error)
	UpdateProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error)
	DeleteProduct(ctx context.Context, req entity.GetProduct) (entity.ResponseHTTP, error)
}
