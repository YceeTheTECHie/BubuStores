package models
type Product struct {
	ProductID int 	`json:"productId"`
	ProductName string 	`json:"productName"`
	ProductImage string	`json:"ProductImage"`
	PricePerUnit string		`json:"pricePerUnit"`
	QuantityAvailable int	`json:"quantityAvailable"`
	Manufacturer string 	`json:"manufacturer"`
	Categories []string 	`json:"category"`
	
	
}

