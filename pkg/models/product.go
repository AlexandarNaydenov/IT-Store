package models

type Product struct {
	ID    string  `json:"id" bson:"_id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
