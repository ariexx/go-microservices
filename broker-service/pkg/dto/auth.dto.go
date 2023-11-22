package dto

type RequestPayload struct {
	Action string       `json:"action"`
	Auth   AuthPayload  `json:"auth,omitempty"`
	Order  OrderPayload `json:"order,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OrderPayload struct {
	PlayerID     string `json:"player_id"`
	ProductID    string `json:"product_id"`
	PaymentCode  string `json:"payment_code"`
	ProductPrice int64  `json:"product_price"`
	EmailAccount string `json:"email_account"`
}
