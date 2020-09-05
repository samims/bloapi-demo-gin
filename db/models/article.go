package models

// Article model
type Article struct {
	ID          int64 `json:"id" orm:"id"`
	Title       string
	Description string
}
