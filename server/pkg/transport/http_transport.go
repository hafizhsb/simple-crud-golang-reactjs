package transport

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"stock/pkg/common"
	"stock/pkg/endpoint"
	"stock/pkg/service"

	"github.com/go-kit/kit/log"

	"stock/pkg/entity"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(s service.Service, logger log.Logger) http.Handler {
	pr := mux.NewRouter()
	e := endpoint.MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(common.SetRequestContext, common.SetUUIDContext),
	}

	r := pr.PathPrefix("/api").Subrouter()

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		e.Status,
		decodeEmpty,
		encodeResponseHTTP,
		options...,
	))

	r.Methods("GET").Path("/products").Handler(httptransport.NewServer(
		e.GetProducts,
		decodeGetProducts,
		encodeResponseHTTP,
		options...,
	))

	r.Methods("GET").Path("/products/{id}").Handler(httptransport.NewServer(
		e.GetProduct,
		decodeGetProduct,
		encodeResponseHTTP,
		options...,
	))

	// r.Methods("POST").Path("/products").Handler(httptransport.NewServer(
	// 	e.CreateProduct,
	// 	decodeCreateProduct,
	// 	encodeResponseHTTP,
	// 	options...,
	// ))

	r.Methods("POST", "OPTIONS").Path("/products").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(httptransport.NewServer(e.CreateProduct, decodeCreateProduct, encodeResponseHTTP, options...)))

	r.Methods("PUT", "OPTIONS").Path("/products/{id}").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedMethods([]string{"PUT", "DELETE"}),
		)(httptransport.NewServer(e.UpdateProduct, decodeUpdateProduct, encodeResponseHTTP, options...)))

	r.Methods("DELETE", "OPTIONS").Path("/products/{id}").Handler(
		handlers.CORS(
			handlers.AllowedMethods([]string{"DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(httptransport.NewServer(e.DeleteProduct, decodeDeleteProduct, encodeResponseHTTP, options...)))

	return pr
}

// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func decodeEmpty(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func decodeGetProducts(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.GetProduct

	return req, nil
}

func decodeGetProduct(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.DetailProduct
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return req, nil
	}

	req.ID = id
	return req, nil
}

func decodeCreateProduct(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.DetailProduct
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeUpdateProduct(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.DetailProduct

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return req, nil
	}

	req.ID = id
	return req, nil
}

func decodeDeleteProduct(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.GetProduct
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return req, nil
	}

	req.ID = id
	return req, nil
}

// Encode Response HTTP
func encodeResponseHTTP(_ context.Context, w http.ResponseWriter, response interface{}) error {
	//TODO: encodeError for buinesss-logic error (optional)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusBadRequest)

	if err == io.EOF {
		err = errors.New(common.MESSAGE_BAD_REQ)
	}
	_ = json.NewEncoder(w).Encode(entity.ResponseHTTP{
		CorrelationID: "",
		Status:        0,
		Message:       err.Error(),
	})
}
