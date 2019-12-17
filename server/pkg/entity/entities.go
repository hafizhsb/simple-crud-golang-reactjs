package entity

type ResponseHTTP struct {
	CorrelationID string      `json:"correlation_id"`
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data,omitempty"`
}

type DetailProduct struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Code        string `json:"code"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
}

type GetProduct struct {
	ID string `json:"id"`
}
