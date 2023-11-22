package dto

type ProductResponse struct {
	ID                    uint32      `json:"id"`
	Name                  string      `json:"name"`
	Banner                string      `json:"banner"`
	ProductDetailResponse interface{} `json:"product_detail"`
}

type ProductDetailResponse struct {
	ID        uint32 `json:"id"`
	ProductId uint32 `json:"product_id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
}
