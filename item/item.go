package item

// Item represents the model of the Item table in the database
type Item struct {
	ID 				int			`json:"id"`
	Manufacturer 	string		`json:"manufacturer"`
	ItemType 		string		`json:"itemType"`
	Price 			float32		`json:"price"`
	Quantity 		int			`json:"quantity"`
}


