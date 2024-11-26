package entities

type ProductEntity struct {
	ProductID string `json:"ProductID"`
	Product string `json:"Product"`
	Price	int	`json:"Price"`
	Description string `json:"Description"`
	Category string `json:"Category"`
}
