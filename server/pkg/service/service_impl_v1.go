package service

import (
	"context"
	"database/sql"
	"fmt"
	"stock/pkg/entity"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
)

type serviceImplV1 struct {
	logger log.Logger
	db     *sql.DB
}

func NewServiceImplV1(logger log.Logger, db *sql.DB) Service {
	return &serviceImplV1{
		logger: logger,
		db:     db,
	}
}

func (s serviceImplV1) Status(ctx context.Context, req interface{}) (entity.ResponseHTTP, error) {
	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "connected",
	}, nil
}

func (s serviceImplV1) GetProducts(ctx context.Context, req entity.GetProduct) (entity.ResponseHTTP, error) {
	var result []entity.DetailProduct
	rows, err := s.db.Query("SELECT id, name, code, price, qty FROM product WHERE is_deleted = 0")
	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
	}

	for rows.Next() {
		each := entity.DetailProduct{}
		err := rows.Scan(&each.ID, &each.ProductName, &each.Code, &each.Price, &each.Qty)
		if err != nil {
			level.Error(s.logger).Log("err", err.Error())
		}

		result = append(result, each)
	}

	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "success",
		Data:          result,
	}, nil
}

func (s serviceImplV1) GetProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error) {
	err := s.db.QueryRow(`SELECT name, code, price, qty FROM product WHERE id = ?`, req.ID).Scan(&req.ProductName, &req.Code, &req.Price, &req.Qty)
	if err != nil {
		level.Error(s.logger).Log("err", err.Error)
		return entity.ResponseHTTP{
			CorrelationID: "",
			Status:        0,
			Message:       err.Error(),
		}, nil
	}

	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "success",
		Data:          req,
	}, nil
}

func (s serviceImplV1) CreateProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error) {

	_, err := s.db.Exec(`INSERT INTO product (name, code, price, qty) VALUES (?, ?, ?, ?)`, req.ProductName, req.Code, req.Price, req.Qty)
	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
		return entity.ResponseHTTP{
			CorrelationID: "",
			Status:        0,
			Message:       err.Error(),
		}, nil
	}
	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "success",
	}, nil
}

func (s serviceImplV1) UpdateProduct(ctx context.Context, req entity.DetailProduct) (entity.ResponseHTTP, error) {
	fmt.Println(req)
	_, err := s.db.Exec(`UPDATE product SET name = ?, code = ?, price = ?, qty = ? WHERE id = ?`, req.ProductName, req.Code, req.Price, req.Qty, req.ID)
	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
		return entity.ResponseHTTP{
			CorrelationID: "",
			Status:        0,
			Message:       err.Error(),
		}, nil
	}
	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "success",
	}, nil
}

func (s serviceImplV1) DeleteProduct(ctx context.Context, req entity.GetProduct) (entity.ResponseHTTP, error) {
	_, err := s.db.Exec(`UPDATE product SET is_deleted = 1 WHERE id = ?`, req.ID)
	if err != nil {
		level.Error(s.logger).Log("err", err.Error())
		return entity.ResponseHTTP{
			CorrelationID: "",
			Status:        0,
			Message:       err.Error(),
		}, nil
	}
	return entity.ResponseHTTP{
		CorrelationID: "",
		Status:        1,
		Message:       "success",
	}, nil
}
