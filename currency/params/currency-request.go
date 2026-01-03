package params

type CurrencyRequestParams struct {
	UserID      uint   `json:"userId"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}
