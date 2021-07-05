package common

type Vegetable struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}
