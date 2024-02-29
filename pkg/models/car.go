package models

type CarModel struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int            `json:"manufacturerId"` // foreign key
	CategoryId     int            `json:"categoryId"`     // foreign key
	Year           int            `json:"year"`
	Specs          Specifications `json:"specifications"` // specification object
	Phot           string         `json:"image"`          // address to photo
}