package models

type Account struct {
	Name   string  `json:"name"`
	Cash   float64 `json:"cash"`
	UserId any     `json:"id"`
}

/*
map := {
	[accountName]: [accountCash]
}
*/