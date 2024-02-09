package dto

type PaymentChannelResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Banner      string `json:"banner"`
	Description string `json:"description"`
}
