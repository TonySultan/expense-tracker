package models

type Account struct {
	Name   string `json:"name"`
	Cash   int    `json:"cash"`
	UserId int    `json:"id"`
}
