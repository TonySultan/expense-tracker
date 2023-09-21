package models

type Transaction struct {
	Amount     float64
	Comment    string
	AccountID  int
	CategoryID int
}
