package models

type Transaction struct {
	Amount     float64 `json:"amount"`
	Comment    string  `json:"comment"`
	AccountID  int     `json:"accountId"`
	CategoryID int     `json:"categoryId"`
}
