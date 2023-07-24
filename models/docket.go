package models

type Docket struct {
	OrderNo       string
	Customer      string
	PickUpPoint   string
	DeliveryPoint string
	Quantity      float64
	Volume        float64
	Status        string
	TruckNo       string
	LogsheetNo    string
	//CreatedAt     string
	//UpdatedAt     string
	//DeletedAt     string
}
