package params

type CreateUserDto struct {
	Name        string
	PhoneNumber string
}

type CurrencyRequestParams struct {
	UserID      uint   `json:"userId"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}
