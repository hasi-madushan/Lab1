package common

type Vegetable struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	UnitPrice    float32 `json:"unitPrice"`
	AvailableQty float32 `json:"availableAmount"`
}
