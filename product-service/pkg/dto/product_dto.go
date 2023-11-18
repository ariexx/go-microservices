package dto

type CreateProductRequest struct {
	ProductName    string                       `json:"product_name"`
	Banner         string                       `json:"banner"`
	ProductDetails []CreateProductDetailRequest `json:"product_details"`
}

type CreateProductDetailRequest struct {
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
}

type ProductResponse struct {
	ID                    uint                    `json:"id"`
	Name                  string                  `json:"name"`
	Banner                string                  `json:"banner"`
	ProductDetailResponse []ProductDetailResponse `json:"product_details"`
}

type ProductDetailResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
}
