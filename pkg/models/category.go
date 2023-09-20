package models

type Category struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	UserId int    `json:"id"`
}
