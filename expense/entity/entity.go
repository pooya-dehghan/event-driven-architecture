package entity

type Expense struct {
	Id     string `json:"id"`
	UserId string `json:"userid"`
	Amount int    `json:"amount"`
	Desc   string `json:"desc"`
}
