package models

type Posts struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
}
